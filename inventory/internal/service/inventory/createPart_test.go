package inventory

import (
	"errors"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stas-tsarev/dz1_order_service/inventory/internal/service/converter"
	"github.com/stas-tsarev/dz1_order_service/inventory/internal/service/model"
)

func (s *ServiceSuite) TestCreatePart_Success() {
	name := gofakeit.Name()
	description := gofakeit.Name()
	price := gofakeit.Price(1.00, 1000000.00)
	stockQuantity := gofakeit.Int64()
	category := model.CATEGORY_ENGINE
	dimensions := model.Dimensions{
		Length: gofakeit.Float64(),
		Width:  gofakeit.Float64(),
		Height: gofakeit.Float64(),
		Weight: gofakeit.Float64(),
	}
	manufacturer := model.Manufacturer{
		Name:    gofakeit.Name(),
		Country: gofakeit.Country(),
		WebSite: gofakeit.Name(),
	}
	tags := make([]string, gofakeit.IntN(10))
	for i := 0; i < len(tags); i++ {
		tags[i] = gofakeit.Word()
	}
	metadata := map[string]model.Value{}
	for i := 0; i < gofakeit.IntN(10); i++ {
		metadata[gofakeit.Word()] = model.Int64Value{Value: int64(gofakeit.Number(1, 1000))}
	}

	partInfo := model.PartInfo{
		Name:          name,
		Description:   description,
		Price:         price,
		StockQuantity: stockQuantity,
		Category:      category,
		Dimensions:    dimensions,
		Manufacturer:  manufacturer,
		Tags:          tags,
		Metadata:      metadata,
	}

	expectedUuid := gofakeit.UUID()
	s.InventoryRepository.On("CreatePart", s.ctx, converter.ServicePartInfoToRepoPartInfo(partInfo)).Return(expectedUuid, nil)

	uuid, err := s.service.CreatePart(s.ctx, partInfo)

	s.Require().NoError(err)
	s.Equal(expectedUuid, uuid)
}

func (s *ServiceSuite) TestCreatePart_FailCreate() {
	partInfo := model.PartInfo{
		Name:          "",
		Description:   "",
		Price:         0,
		StockQuantity: 0,
		Category:      0,
		Dimensions:    model.Dimensions{},
		Manufacturer:  model.Manufacturer{},
		Tags:          nil,
		Metadata:      nil,
	}

	expectedError := errors.New("create part error")
	s.InventoryRepository.On("CreatePart", s.ctx, converter.ServicePartInfoToRepoPartInfo(partInfo)).Return("", expectedError)

	uuid, err := s.service.CreatePart(s.ctx, partInfo)

	s.Equal("", uuid)
	s.Error(err)
}
