package order

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stas-tsarev/dz1_order_service/order/internal/errs"
	modelRepo "github.com/stas-tsarev/dz1_order_service/order/internal/repository/model"
	"github.com/stas-tsarev/dz1_order_service/order/internal/service/converter"
)

func (s *ServiceSuite) TestGetOrder_Success() {
	partUuids := make([]string, gofakeit.Number(1, 10))
	for i := 0; i < len(partUuids); i++ {
		partUuids[i] = gofakeit.UUID()
	}

	var (
		orderUuid  = gofakeit.UUID()
		userUuid   = gofakeit.UUID()
		totalPrice = gofakeit.Float64()
		status     = "PENDING_PAYMENT"

		orderRepo = modelRepo.Order{
			OrderUuid:       orderUuid,
			UserUuid:        userUuid,
			PartUuids:       partUuids,
			TotalPrice:      totalPrice,
			TransactionUuid: nil,
			PaymentMethod:   nil,
			Status:          status,
		}
	)

	s.orderRepository.On("GetOrder", s.ctx, orderUuid).Return(orderRepo, nil).Once()
	orderService := converter.RepoOrderToServiceOrder(orderRepo)

	res, err := s.service.GetOrder(s.ctx, orderUuid)
	s.NoError(err)
	s.NotEmpty(orderService)
	s.Equal(res, orderService)
}

func (s *ServiceSuite) TestGetOrder_NotFound() {
	var (
		err       = errs.ErrorOrderNotFound
		orderUuid = gofakeit.UUID()
	)

	s.orderRepository.On("GetOrder", s.ctx, orderUuid).Return(modelRepo.Order{}, err).Once()

	res, err2 := s.service.GetOrder(s.ctx, orderUuid)
	s.Error(err2)
	s.ErrorIs(err, err2)
	s.Empty(res)
}
