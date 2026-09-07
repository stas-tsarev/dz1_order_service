package inventory

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stas-tsarev/dz1_order_service/inventory/internal/repository/model"
	"github.com/stas-tsarev/dz1_order_service/inventory/internal/service/converter"
)

func (s *ServiceSuite) TestGetPart_Success() {
	partUuid := gofakeit.UUID()

	part := model.Part{
		Uuid:      "",
		Info:      model.PartInfo{},
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	s.InventoryRepository.On("GetPart", s.ctx, partUuid).Return(part, nil)

	part1, err := s.service.GetPart(s.ctx, partUuid)

	s.Require().NoError(err)
	s.Require().Equal(converter.RepoPartToServicePart(part), part1)
}

func (s *ServiceSuite) TestGetPart_NotFound() {
	partUuid := gofakeit.UUID()

	expectedError := model.ErrorPartNotFound
	s.InventoryRepository.On("GetPart", s.ctx, partUuid).Return(model.Part{}, expectedError)

	part, err := s.service.GetPart(s.ctx, partUuid)

	s.Require().Equal(expectedError, err)
	s.Require().Empty(part)
	s.Require().Error(err)
}
