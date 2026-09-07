package order

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stas-tsarev/dz1_order_service/order/internal/errs"
)

func (s *ServiceSuite) TestCancelOrder_Success() {
	orderUuid := gofakeit.UUID()

	s.orderRepository.On("CancelOrder", s.ctx, orderUuid).Return(nil).Once()

	err := s.service.CancelOrder(s.ctx, orderUuid)

	s.NoError(err)
}

func (s *ServiceSuite) TestCancelOrder_NotFound() {
	orderUuid := gofakeit.UUID()
	expectedError := errs.ErrorOrderNotFound

	s.orderRepository.On("CancelOrder", s.ctx, orderUuid).Return(expectedError).Once()
	err := s.service.CancelOrder(s.ctx, orderUuid)

	s.Error(err)
	s.Equal(expectedError, err)
}

func (s *ServiceSuite) TestCancelOrder_Conflict() {
	orderUuid := gofakeit.UUID()
	expectedError := errs.ErrorConflict

	s.orderRepository.On("CancelOrder", s.ctx, orderUuid).Return(expectedError).Once()
	err := s.service.CancelOrder(s.ctx, orderUuid)

	s.Error(err)
	s.Equal(expectedError, err)
}
