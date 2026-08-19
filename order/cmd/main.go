package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	inventory_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/proto/inventory/v1"
	payment_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/proto/payment/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	order_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/openapi/order/v1"
)

const (
	httpPortOrder     = "8080"
	readHeaderTimeout = 5 * time.Second
	shutdownTimeout   = 10 * time.Second
)

type OrderStorage struct {
	mu    sync.Mutex
	order map[string]*order_v1.Order
}

func NewOrderStorage() *OrderStorage {
	return &OrderStorage{
		order: make(map[string]*order_v1.Order),
	}
}

type Clients struct {
	clientInv inventory_v1.InventoryServiceClient
	clientPay payment_v1.PaymentServiceClient
}

// OrderHandler реализует интерфейс order_v1.Handler для обработки запросов к API заказов
type OrderHandler struct {
	storage *OrderStorage
	clients *Clients
}

// NewOrderHandler создает новый обработчик запросов к API заказов
func NewOrderHandler(storage *OrderStorage, clients *Clients) *OrderHandler {
	return &OrderHandler{
		storage: storage,
		clients: clients,
	}
}

func (h *OrderHandler) CancelOrder(_ context.Context, req order_v1.OptDeleteOrderRequest, params order_v1.CancelOrderParams) (order_v1.CancelOrderRes, error) {
	return nil, nil
}

// TODO эту хуйню полностью надо переделать, еще обратить внимание на price, который float32, когда в inventory он float64
func (h *OrderHandler) CreateOrder(_ context.Context, req order_v1.OptCreateOrderRequest) (order_v1.CreateOrderRes, error) {
	h.storage.mu.Lock()
	defer h.storage.mu.Unlock()

	request := req.Value

	partsUUIDs := make([]string, 0, 0)
	for _, part := range request.PartUuids {
		partsUUIDs = append(partsUUIDs, part.String())
	}
	listPartReq := &inventory_v1.ListPartsRequest{
		Filter: &inventory_v1.PartsFilter{
			Uuids:                 partsUUIDs,
			Names:                 nil,
			Categories:            nil,
			ManufacturerCountries: nil,
			Tags:                  nil,
		},
	}

	listPartResp, err := h.clients.clientInv.ListParts(context.Background(), listPartReq)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	foundParts := make(map[string]bool)
	for _, part := range listPartResp.Parts {
		foundParts[part.Uuid] = true
	}

	var missingParts []string
	for _, part := range request.PartUuids {
		if !foundParts[part.String()] {
			missingParts = append(missingParts, part.String())
		}
	}

	if len(missingParts) > 0 {
		return nil, errors.New(strings.Join(missingParts, ","))
	}

	var totalPrice float32
	for _, part := range listPartResp.Parts {
		totalPrice += float32(part.Info.Price)
	}

	orderUuid := uuid.New()

	order := &order_v1.Order{
		OrderUUID:  orderUuid,
		UserUUID:   request.UserUUID,
		PartUuids:  request.PartUuids,
		TotalPrice: totalPrice,
		Status:     "PENDIND_PAYMENT",
	}

	h.storage.order[orderUuid.String()] = order
	log.Printf("Order %s created", orderUuid.String())

	return &order_v1.CreateOrderResponse{
		orderUuid,
		totalPrice,
	}, nil
}

func (h *OrderHandler) GetOrder(_ context.Context, params order_v1.GetOrderParams) (order_v1.GetOrderRes, error) {
	return nil, nil
}

func (h *OrderHandler) PayOrder(_ context.Context, req order_v1.OptPayOrderRequest, params order_v1.PayOrderParams) (order_v1.PayOrderRes, error) {
	return nil, nil
}

func main() {
	// Добавляем клиентов
	// Inventory Client
	ctx := context.Background()
	connInv, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := connInv.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	clientInv := inventory_v1.NewInventoryServiceClient(connInv)

	// Payment Client
	connPay, err := grpc.NewClient(
		"localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := connPay.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	clientPay := payment_v1.NewPaymentServiceClient(connPay)

	clients := &Clients{clientInv: clientInv, clientPay: clientPay}

	// Создание хранилища для заказов
	storage := NewOrderStorage()

	orderHandler := NewOrderHandler(storage, clients)

	orderServer, err := order_v1.NewServer(orderHandler)
	if err != nil {
		log.Fatalf("Error creating order server: %v", err)
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))

	r.Mount("/", orderServer)

	server := &http.Server{
		Addr:              net.JoinHostPort("localhost", httpPortOrder),
		Handler:           r,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	go func() {
		log.Printf("🚀 HTTP-сервер запущен на порту %s\n", httpPortOrder)
		err = server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("❌ Ошибка запуска сервера: %v\n", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Завершение работы сервера...")

	// Создаем контекст с таймаутом для остановки сервера
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		log.Printf("❌ Ошибка при остановке сервера: %v\n", err)
	}

	log.Println("✅ Сервер остановлен")
}
