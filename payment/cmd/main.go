package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	paymentApi "github.com/stas-tsarev/dz1_order_service/payment/internal/api/payment/v1"
	paymentService "github.com/stas-tsarev/dz1_order_service/payment/internal/service/payment"
	payment_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/proto/payment/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPortPayment = 50052

//
//type PaymentService struct {
//	payment_v1.UnimplementedPaymentServiceServer
//	mu sync.Mutex
//}
//
//func (s *PaymentService) PayOrder(_ context.Context, req *payment_v1.PayOrderRequest) (*payment_v1.PayOrderResponse, error) {
//	s.mu.Lock()
//	defer s.mu.Unlock()
//
//	if req.GetPay() == nil {
//		return nil, status.Errorf(codes.InvalidArgument, "missing payment information")
//	}
//
//	newUUID := uuid.NewString()
//
//	log.Printf("Create payment with uuid: %s", newUUID)
//
//	return &payment_v1.PayOrderResponse{
//		TransactionUuid: newUUID,
//	}, nil
//}
//
//func main() {
//	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPortPayment))
//
//	if err != nil {
//		log.Printf("failed to listen: %v\n", err)
//		return
//	}
//	//defer func() {
//	//	if cerr := lis.Close(); cerr != nil {
//	//		log.Printf("failed to close listener: %v\n", cerr)
//	//	}
//	//}()
//
//	// Создание сервера
//	s := grpc.NewServer()
//
//	// Регистрация сервиса
//	service := &PaymentService{}
//
//	payment_v1.RegisterPaymentServiceServer(s, service)
//
//	// Включение рефлексии для отладки
//	reflection.Register(s)
//
//	go func() {
//		log.Printf("🤯 gRPC server listening on %d\n", grpcPortPayment)
//		err = s.Serve(lis)
//		if err != nil {
//			log.Printf("failed to serve: %v\n", err)
//			return
//		}
//	}()
//
//	quit := make(chan os.Signal, 1)
//	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
//	<-quit
//	log.Println("Shutting down server...")
//	s.GracefulStop()
//	lis.Close()
//	log.Println("Server stopped")
//}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPortPayment))
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
		return
	}

	// Создание сервера
	s := grpc.NewServer()

	service := paymentService.NewService()
	api := paymentApi.NewApi(service)

	payment_v1.RegisterPaymentServiceServer(s, api)

	// Включение рефлексии для отладки
	reflection.Register(s)

	go func() {
		log.Printf("🤯 gRPC server listening on %d\n", grpcPortPayment)
		err = s.Serve(lis)
		if err != nil {
			log.Printf("failed to serve: %v\n", err)
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	s.GracefulStop()
	lis.Close()
	log.Println("Server stopped")
}
