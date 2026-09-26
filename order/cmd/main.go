package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	orderApi "github.com/stas-tsarev/dz1_order_service/order/internal/api/order/v1"
	inventoryclient "github.com/stas-tsarev/dz1_order_service/order/internal/client/grpc/inventory/v1"
	paymentclient "github.com/stas-tsarev/dz1_order_service/order/internal/client/grpc/payment/v1"
	orderRepo "github.com/stas-tsarev/dz1_order_service/order/internal/repository/order"
	orderService "github.com/stas-tsarev/dz1_order_service/order/internal/service/order"
	order_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/openapi/order/v1"
	inventory_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/proto/inventory/v1"
	payment_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/proto/payment/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	httpPortOrder     = "8080"
	readHeaderTimeout = 5 * time.Second
	shutdownTimeout   = 10 * time.Second
)

//type OrderStorage struct {
//	mu    sync.Mutex
//	order map[string]*order_v1.Order
//}
//
//func NewOrderStorage() *OrderStorage {
//	return &OrderStorage{
//		order: make(map[string]*order_v1.Order),
//	}
//}
//
//type Clients struct {
//	clientInv inventory_v1.InventoryServiceClient
//	clientPay payment_v1.PaymentServiceClient
//}
//
//// OrderHandler реализует интерфейс order_v1.Handler для обработки запросов к API заказов
//type OrderHandler struct {
//	storage *OrderStorage
//	clients *Clients
//}
//
//// NewOrderHandler создает новый обработчик запросов к API заказов
//func NewOrderHandler(storage *OrderStorage, clients *Clients) *OrderHandler {
//	return &OrderHandler{
//		storage: storage,
//		clients: clients,
//	}
//}
//
//func (h *OrderHandler) CancelOrder(_ context.Context, params order_v1.CancelOrderParams) (order_v1.CancelOrderRes, error) {
//	h.storage.mu.Lock()
//	defer h.storage.mu.Unlock()
//
//	order := h.storage.order[params.OrderUUID.String()]
//	if order == nil {
//		err := errors.New("Order not found")
//		log.Println("CancelOrder: ", err)
//		return &order_v1.R404NotFound{
//			Code:    404,
//			Message: err.Error(),
//		}, nil
//	}
//
//	if order.Status == order_v1.OrderStatusPAID {
//		err := errors.New("Cannot cancel order because it is paid")
//		log.Println("CancelOrder: ", err)
//		return &order_v1.R409Conflict{
//			Code:    409,
//			Message: err.Error(),
//		}, nil
//	}
//
//	canceledOrder := &order_v1.Order{
//		OrderUUID:     order.OrderUUID,
//		UserUUID:      order.UserUUID,
//		PartUuids:     order.PartUuids,
//		TotalPrice:    order.TotalPrice,
//		PaymentMethod: order.PaymentMethod,
//		Status:        order_v1.OrderStatusCANCELED,
//	}
//
//	h.storage.order[params.OrderUUID.String()] = canceledOrder
//
//	return &order_v1.CancelOrderNoContent{}, nil
//}
//
//func (h *OrderHandler) CreateOrder(_ context.Context, req order_v1.OptCreateOrderRequest) (order_v1.CreateOrderRes, error) {
//	h.storage.mu.Lock()
//	defer h.storage.mu.Unlock()
//
//	request := req.Value
//
//	// Сохраняем детали для поиска в отдельную переменную
//	partsUUIDs := make([]string, 0, 0)
//	for _, part := range request.PartUuids {
//		partsUUIDs = append(partsUUIDs, part.String())
//	}
//
//	// Создаем запрос для поиска деталей в order
//	listPartReq := &inventory_v1.ListPartsRequest{
//		Filter: &inventory_v1.PartsFilter{
//			Uuids:                 partsUUIDs,
//			Names:                 nil,
//			Categories:            nil,
//			ManufacturerCountries: nil,
//			Tags:                  nil,
//		},
//	}
//
//	listPartResp, _ := h.clients.clientInv.ListParts(context.Background(), listPartReq)
//	if listPartResp == nil {
//		err := errors.New("Parts not found in order")
//		log.Println("CreateOrder: ", err)
//		return &order_v1.R422UnprocessableEntity{
//			Code:    422,
//			Message: err.Error(),
//		}, nil
//	}
//	if len(listPartResp.Parts) != len(partsUUIDs) {
//		err := errors.New("Some parts not found in order")
//		log.Println("CreateOrder: ", err)
//		return &order_v1.R422UnprocessableEntity{
//			Code:    422,
//			Message: err.Error(),
//		}, nil
//	}
//
//	foundParts := make(map[string]bool)
//	for _, part := range listPartResp.Parts {
//		foundParts[part.Uuid] = true
//	}
//
//	var totalPrice float64
//	for _, part := range listPartResp.Parts {
//		totalPrice += part.Info.Price
//	}
//
//	orderUuid := uuid.New()
//
//	order := &order_v1.Order{
//		OrderUUID:  orderUuid,
//		UserUUID:   request.UserUUID,
//		PartUuids:  request.PartUuids,
//		TotalPrice: totalPrice,
//		Status:     order_v1.OrderStatusPENDINGPAYMENT,
//	}
//
//	h.storage.order[orderUuid.String()] = order
//	log.Printf("Order %s created", orderUuid.String())
//
//	return &order_v1.CreateOrderResponse{
//		OrderUUID:  orderUuid,
//		TotalPrice: totalPrice,
//	}, nil
//}
//
//func (h *OrderHandler) GetOrder(_ context.Context, params order_v1.GetOrderParams) (order_v1.GetOrderRes, error) {
//	h.storage.mu.Lock()
//	defer h.storage.mu.Unlock()
//
//	order := h.storage.order[params.OrderUUID.String()]
//	if order == nil {
//		err := errors.New("Order not found")
//		log.Println("GetOrder: ", err)
//		return &order_v1.R404NotFound{
//			Code:    404,
//			Message: err.Error(),
//		}, nil
//	}
//
//	return &order_v1.Order{
//		OrderUUID:       order.OrderUUID,
//		UserUUID:        order.UserUUID,
//		PartUuids:       order.PartUuids,
//		TotalPrice:      order.TotalPrice,
//		TransactionUUID: order.TransactionUUID,
//		PaymentMethod:   order.PaymentMethod,
//		Status:          order.Status,
//	}, nil
//}
//
//// Функция переводит order.PaymentMethod в payment.PaymentMethod
//func GetPaymentMethod(method order_v1.PaymentMethod) payment_v1.PaymentMethod {
//	switch method {
//	case order_v1.PaymentMethodUNKNOWN:
//		return payment_v1.PaymentMethod_PAYMENT_METHOD_UNKNOWN_UNSPECIFIED
//	case order_v1.PaymentMethodCARD:
//		return payment_v1.PaymentMethod_PAYMENT_METHOD_CARD
//	case order_v1.PaymentMethodSBP:
//		return payment_v1.PaymentMethod_PAYMENT_METHOD_SBP
//	case order_v1.PaymentMethodCREDITCARD:
//		return payment_v1.PaymentMethod_PAYMENT_METHOD_CREDIT_CARD
//	case order_v1.PaymentMethodINVESTORMONEY:
//		return payment_v1.PaymentMethod_PAYMENT_METHOD_INVESTOR_MONEY
//	}
//	return payment_v1.PaymentMethod_PAYMENT_METHOD_UNKNOWN_UNSPECIFIED
//}
//
//func (h *OrderHandler) PayOrder(_ context.Context, req order_v1.OptPayOrderRequest, params order_v1.PayOrderParams) (order_v1.PayOrderRes, error) {
//	h.storage.mu.Lock()
//	defer h.storage.mu.Unlock()
//
//	order := h.storage.order[params.OrderUUID.String()]
//	if order == nil {
//		err := errors.New("Order not found")
//		log.Println("PayOrder: ", err)
//		return &order_v1.R404NotFound{
//			Code:    404,
//			Message: err.Error(),
//		}, nil
//	}
//
//	payOrder := &payment_v1.PayOrderRequest{
//		Pay: &payment_v1.Pay{
//			OrderUuid:     order.OrderUUID.String(),
//			UserUuid:      order.UserUUID.String(),
//			PaymentMethod: GetPaymentMethod(req.Value.PaymentMethod),
//		},
//	}
//
//	payment, _ := h.clients.clientPay.PayOrder(context.Background(), payOrder)
//	transactionUuid, _ := uuid.Parse(payment.TransactionUuid)
//
//	updatedOrder := &order_v1.Order{
//		OrderUUID:  order.OrderUUID,
//		UserUUID:   order.UserUUID,
//		PartUuids:  order.PartUuids,
//		TotalPrice: order.TotalPrice,
//		TransactionUUID: order_v1.NilUUID{
//			Value: transactionUuid,
//			Null:  false,
//		},
//		Status:        order_v1.OrderStatusPAID,
//		PaymentMethod: req.Value.PaymentMethod,
//	}
//
//	log.Printf("Order %s payed", updatedOrder.OrderUUID)
//	h.storage.order[params.OrderUUID.String()] = updatedOrder
//
//	return &order_v1.PayOrderResponse{
//		TransactionUUID: transactionUuid,
//	}, nil
//}
//
//func main() {
//	// Добавляем клиентов
//	// Inventory Client
//	ctx := context.Background()
//	connInv, err := grpc.NewClient(
//		"localhost:50051",
//		grpc.WithTransportCredentials(insecure.NewCredentials()),
//	)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer func() {
//		if err := connInv.Close(); err != nil {
//			log.Fatal(err)
//		}
//	}()
//	clientInv := inventory_v1.NewInventoryServiceClient(connInv)
//
//	// Payment Client
//	connPay, err := grpc.NewClient(
//		"localhost:50052",
//		grpc.WithTransportCredentials(insecure.NewCredentials()),
//	)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer func() {
//		if err := connPay.Close(); err != nil {
//			log.Fatal(err)
//		}
//	}()
//	clientPay := payment_v1.NewPaymentServiceClient(connPay)
//
//	clients := &Clients{clientInv: clientInv, clientPay: clientPay}
//
//	// Создание хранилища для заказов
//	storage := NewOrderStorage()
//
//	orderHandler := NewOrderHandler(storage, clients)
//
//	orderServer, err := order_v1.NewServer(orderHandler)
//	if err != nil {
//		log.Fatalf("Error creating order server: %v", err)
//	}
//
//	r := chi.NewRouter()
//
//	r.Use(middleware.Logger)
//	r.Use(middleware.Recoverer)
//	r.Use(middleware.Timeout(10 * time.Second))
//
//	r.Mount("/", orderServer)
//
//	server := &http.Server{
//		Addr:              net.JoinHostPort("localhost", httpPortOrder),
//		Handler:           r,
//		ReadHeaderTimeout: readHeaderTimeout,
//	}
//
//	go func() {
//		log.Printf("🚀 HTTP-сервер запущен на порту %s\n", httpPortOrder)
//		err = server.ListenAndServe()
//		if err != nil && !errors.Is(err, http.ErrServerClosed) {
//			log.Printf("❌ Ошибка запуска сервера: %v\n", err)
//		}
//	}()
//
//	// Graceful shutdown
//	quit := make(chan os.Signal, 1)
//	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
//	<-quit
//
//	log.Println("🛑 Завершение работы сервера...")
//
//	// Создаем контекст с таймаутом для остановки сервера
//	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
//	defer cancel()
//
//	err = server.Shutdown(ctx)
//	if err != nil {
//		log.Printf("❌ Ошибка при остановке сервера: %v\n", err)
//	}
//
//	log.Println("✅ Сервер остановлен")
//}

func main() {
	ctx := context.Background()

	connInv, err := grpc.NewClient(
		"inventory:50051",
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
		"payment:50052",
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

	generatedClientPay := paymentclient.NewClient(clientPay)
	generatedClientInv := inventoryclient.NewClient(clientInv)
	repo := orderRepo.NewRepository()
	service := orderService.NewService(repo, generatedClientInv, generatedClientPay)
	api := orderApi.NewApi(service)

	orderServer, err := order_v1.NewServer(api)
	if err != nil {
		log.Fatalf("Error creating order server: %v", err)
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))

	r.Mount("/", orderServer)

	server := &http.Server{
		Addr:              net.JoinHostPort("0.0.0.0", httpPortOrder),
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
