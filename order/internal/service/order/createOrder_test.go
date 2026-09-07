package order

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stas-tsarev/dz1_order_service/order/internal/client/model"
	"github.com/stas-tsarev/dz1_order_service/order/internal/errs"
)

func (s *ServiceSuite) TestCreateOrder_Success() {
	userUuid := gofakeit.UUID()
	partUuids := make([]string, 3)
	for i := 0; i < len(partUuids); i++ {
		partUuids[i] = gofakeit.UUID()
	}

	expectedParts := []model.Part{
		{
			Uuid: partUuids[0],
			Info: model.PartInfo{
				Price: 100.50,
			},
			CreatedAt: gofakeit.Date(),
			UpdatedAt: gofakeit.Date(),
		},
		{
			Uuid: partUuids[1],
			Info: model.PartInfo{
				Price: 120.00,
			},
			CreatedAt: gofakeit.Date(),
			UpdatedAt: gofakeit.Date(),
		},
		{
			Uuid: partUuids[2],
			Info: model.PartInfo{
				Price: 130.00,
			},
			CreatedAt: gofakeit.Date(),
			UpdatedAt: gofakeit.Date(),
		},
	}
	totalPrice := 100.50 + 120.00 + 130.00

	s.inventoryClient.On("ListParts", s.ctx, model.PartsFilter{Uuids: partUuids}).Return(expectedParts, nil)

	expectedOrderUuid := gofakeit.UUID()
	s.orderRepository.On("CreateOrder", s.ctx, userUuid, partUuids, totalPrice).Return(expectedOrderUuid, totalPrice, nil)

	uuid, price, err := s.service.CreateOrder(s.ctx, userUuid, partUuids)
	s.Require().NoError(err)
	s.Require().Equal(expectedOrderUuid, uuid)
	s.Require().Equal(price, totalPrice)
}

func (s *ServiceSuite) TestCreateOrder_UnprocessableEntity() {
	userUuid := gofakeit.UUID()
	partUuids := make([]string, 3)
	for i := 0; i < len(partUuids); i++ {
		partUuids[i] = gofakeit.UUID()
	}

	// Возвращаем ТОЛЬКО 2 детали (третья отсутствует)
	expectedParts := []model.Part{
		{
			Uuid: partUuids[0],
			Info: model.PartInfo{
				Price: 100.50,
			},
			CreatedAt: gofakeit.Date(),
			UpdatedAt: gofakeit.Date(),
		},
		{
			Uuid: partUuids[1],
			Info: model.PartInfo{
				Price: 120.00,
			},
			CreatedAt: gofakeit.Date(),
			UpdatedAt: gofakeit.Date(),
		},
		// partUuids[2] отсутствует
	}

	s.inventoryClient.On("ListParts", s.ctx, model.PartsFilter{Uuids: partUuids}).
		Return(expectedParts, nil)

	// НЕ НАСТРАИВАЕМ мок репозитория, потому что код вернет ошибку ДО вызова

	orderUuid, total, err := s.service.CreateOrder(s.ctx, userUuid, partUuids)

	// Ожидаем ошибку о ненайденной детали
	s.Error(err)
	s.Contains(err.Error(), errs.ErrorUnprocessableEntity.Error())
	s.Empty(orderUuid)
	s.Equal(0.0, total)

	// Проверяем, что репозиторий НЕ вызывался
	s.orderRepository.AssertNotCalled(s.T(), "CreateOrder")
}
