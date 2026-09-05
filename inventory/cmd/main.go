package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	invV1API "github.com/stas-tsarev/dz1_order_service/inventory/internal/api/inventory/v1"
	invRepository "github.com/stas-tsarev/dz1_order_service/inventory/internal/repository/inventory"
	invService "github.com/stas-tsarev/dz1_order_service/inventory/internal/service/inventory"
	inventory_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/proto/inventory/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPortInventory = 50051

//type InventoryService struct {
//	inventory_v1.UnimplementedInventoryServiceServer
//	mu    sync.Mutex
//	parts map[string]*inventory_v1.Part
//}
//
//func (s *InventoryService) GetPart(_ context.Context, req *inventory_v1.GetPartRequest) (*inventory_v1.GetPartResponse, error) {
//	s.mu.Lock()
//	defer s.mu.Unlock()
//
//	part, ok := s.parts[req.GetUuid()]
//	if !ok {
//		return nil, status.Errorf(codes.NotFound, "part: %s", req.GetUuid())
//	}
//
//	return &inventory_v1.GetPartResponse{
//		Part: part,
//	}, nil
//}
//
//// Ниже приведены функции для фильтрации
//
//func (s *InventoryService) ListParts(_ context.Context, req *inventory_v1.ListPartsRequest) (*inventory_v1.ListPartsResponse, error) {
//	s.mu.Lock()
//	defer s.mu.Unlock()
//
//	// Получаем фильтр
//	filter := req.GetFilter()
//
//	// Создаем массив для хранения деталей
//	partsArray := make([]*inventory_v1.Part, 0, len(s.parts))
//	for _, part := range s.parts {
//		partsArray = append(partsArray, part)
//	}
//
//	// Возвращаем все детали, если фильтр пуст
//	if filter == nil {
//		return &inventory_v1.ListPartsResponse{
//			Parts: partsArray,
//		}, nil
//	}
//	if filter.Uuids == nil && filter.Names == nil && filter.Categories == nil && filter.ManufacturerCountries == nil && filter.Tags == nil {
//		return &inventory_v1.ListPartsResponse{
//			Parts: partsArray,
//		}, nil
//	}
//
//	// Фильтруем по UUID
//	if filter.Uuids != nil && len(partsArray) != 0 {
//		uuidSet := make(map[string]bool, len(filter.Uuids))
//		for _, uuid := range filter.Uuids {
//			uuidSet[uuid] = true
//		}
//		result := make([]*inventory_v1.Part, 0, len(partsArray))
//		for _, part := range partsArray {
//			if uuidSet[part.Uuid] {
//				result = append(result, part)
//			}
//		}
//		partsArray = result
//	}
//
//	// Фильтруем по имени
//	if filter.Names != nil && len(partsArray) != 0 {
//		nameSet := make(map[string]bool, len(filter.Names))
//		for _, name := range filter.Names {
//			nameSet[name] = true
//		}
//		result := make([]*inventory_v1.Part, 0, len(partsArray))
//		for _, part := range partsArray {
//			if nameSet[part.Info.Name] {
//				result = append(result, part)
//			}
//		}
//		partsArray = result
//	}
//
//	// Фильтруем по категории
//	if filter.Categories != nil && len(partsArray) != 0 {
//		categorySet := make(map[string]bool, len(filter.Categories))
//		for _, category := range filter.Categories {
//			categorySet[string(category)] = true
//		}
//		result := make([]*inventory_v1.Part, 0, len(partsArray))
//		for _, part := range partsArray {
//			if categorySet[string(part.Info.Category)] {
//				result = append(result, part)
//			}
//		}
//		partsArray = result
//	}
//
//	// Фильтруем по странам производителей
//	if filter.ManufacturerCountries != nil && len(partsArray) != 0 {
//		countrySet := make(map[string]bool, len(filter.ManufacturerCountries))
//		for _, country := range filter.ManufacturerCountries {
//			countrySet[country] = true
//		}
//		result := make([]*inventory_v1.Part, 0, len(partsArray))
//		for _, part := range partsArray {
//			if countrySet[part.Info.Manufacturer.Country] {
//				result = append(result, part)
//			}
//		}
//		partsArray = result
//	}
//
//	// Фильтруем по тегам
//	if filter.Tags != nil && len(partsArray) != 0 {
//		tagSet := make(map[string]bool, len(filter.Tags))
//		for _, tag := range filter.Tags {
//			tagSet[tag] = true
//		}
//		result := make([]*inventory_v1.Part, 0, len(partsArray))
//		for _, part := range partsArray {
//			for ind := range part.Info.Tags {
//				if tagSet[part.Info.Tags[ind]] {
//					result = append(result, part)
//					break
//				}
//			}
//		}
//		partsArray = result
//	}
//
//	// Если массив пуст, то возвращаем ошибку
//	if len(partsArray) == 0 {
//		return nil, status.Errorf(codes.NotFound, "can't found parts with your filter")
//	}
//
//	return &inventory_v1.ListPartsResponse{
//		Parts: partsArray,
//	}, nil
//}
//
//func (s *InventoryService) CreatePart(_ context.Context, req *inventory_v1.CreatePartRequest) (*inventory_v1.CreatePartResponse, error) {
//	s.mu.Lock()
//	defer s.mu.Unlock()
//
//	if req.GetInfo() == nil {
//		return nil, status.Errorf(codes.InvalidArgument, "missing info")
//	}
//
//	newUUID := uuid.NewString()
//
//	goTime := time.Now()
//	protoTimestamp := timestamppb.New(goTime)
//
//	part := &inventory_v1.Part{
//		Uuid:      newUUID,
//		Info:      req.GetInfo(),
//		CreatedAt: protoTimestamp,
//		UpdatedAt: protoTimestamp,
//	}
//
//	s.parts[newUUID] = part
//	log.Printf("Create part with uuid: %s", newUUID)
//
//	return &inventory_v1.CreatePartResponse{
//		Uuid: newUUID,
//	}, nil
//}
//
//func main() {
//	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPortInventory))
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
//	service := &InventoryService{
//		parts: make(map[string]*inventory_v1.Part),
//	}
//
//	inventory_v1.RegisterInventoryServiceServer(s, service)
//
//	// Включение рефлексии для отладки
//	reflection.Register(s)
//
//	go func() {
//		log.Printf("🤯 gRPC server listening on %d\n", grpcPortInventory)
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
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPortInventory))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()

	repo := invRepository.NewRepository()
	service := invService.NewService(repo)
	api := invV1API.NewApi(service)

	inventory_v1.RegisterInventoryServiceServer(s, api)

	reflection.Register(s)

	go func() {
		log.Printf("🤯 gRPC server listening on %d\n", grpcPortInventory)
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
