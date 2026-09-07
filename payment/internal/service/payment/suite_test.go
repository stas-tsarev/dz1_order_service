package payment

import (
	"context"
	"testing"

	"github.com/stas-tsarev/dz1_order_service/payment/internal/service/mocks"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite

	ctx context.Context

	mockService *mocks.PaymentService
	service     *service
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()

	s.mockService = mocks.NewPaymentService(s.T())
	s.service = NewService()
}

func (s *ServiceSuite) TearDownTest() {
	s.mockService.AssertExpectations(s.T())
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
