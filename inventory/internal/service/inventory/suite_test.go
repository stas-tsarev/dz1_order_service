package inventory

import (
	"context"
	"testing"

	"github.com/stas-tsarev/dz1_order_service/inventory/internal/repository/mocks"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite

	ctx context.Context

	InventoryRepository *mocks.InventoryRepository

	service *service
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()

	s.InventoryRepository = mocks.NewInventoryRepository(s.T())

	s.service = NewService(
		s.InventoryRepository,
	)
}

func (s *ServiceSuite) TearDownTest() {

}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
