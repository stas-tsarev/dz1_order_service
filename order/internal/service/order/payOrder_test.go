package order

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-faster/errors"
	"github.com/stas-tsarev/dz1_order_service/order/internal/errs"
)

func (s *ServiceSuite) TestPayOrder_Success() {
	orderUuid := gofakeit.UUID()
	userUuid := gofakeit.UUID()
	paymentMethod := "PAYMENT_METHOD_CARD"

	expectedTransactionUuid := gofakeit.UUID()

	s.paymentClient.On("PayOrder", s.ctx, orderUuid, userUuid, paymentMethod).Return(expectedTransactionUuid, nil).Once()

	s.orderRepository.On("PayOrder", s.ctx, orderUuid, paymentMethod, expectedTransactionUuid).Return(nil).Once()

	transaction, err := s.service.PayOrder(s.ctx, orderUuid, userUuid, paymentMethod)

	s.Require().NoError(err)
	s.Require().Equal(expectedTransactionUuid, transaction)
}

func (s *ServiceSuite) TestPayOrder_NotFound() {
	orderUuid := gofakeit.UUID()
	userUuid := gofakeit.UUID()
	paymentMethod := "PAYMENT_METHOD_SBP"

	expectedTransactionUuid := gofakeit.UUID()
	expectedError := errs.ErrorOrderNotFound

	s.paymentClient.On("PayOrder", s.ctx, orderUuid, userUuid, paymentMethod).Return(expectedTransactionUuid, nil).Once()

	s.orderRepository.On("PayOrder", s.ctx, orderUuid, paymentMethod, expectedTransactionUuid).Return(expectedError).Once()

	transaction, err := s.service.PayOrder(s.ctx, orderUuid, userUuid, paymentMethod)

	s.Require().Error(err)
	s.Require().Equal(expectedError, err)
	s.Require().Equal(transaction, "")
}

func (s *ServiceSuite) TestPayOrder_Error() {
	orderUuid := gofakeit.UUID()
	userUuid := gofakeit.UUID()
	paymentMethod := "PAYMENT_METHOD_CREDIT_CARD"

	expectedError := errors.New("aaa")

	s.paymentClient.On("PayOrder", s.ctx, orderUuid, userUuid, paymentMethod).Return("", expectedError).Once()

	_, err := s.service.PayOrder(s.ctx, orderUuid, userUuid, paymentMethod)

	s.Require().Error(err)
	s.Require().Equal(expectedError, err)
}
