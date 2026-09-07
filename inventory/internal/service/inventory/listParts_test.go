package inventory

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
	model2 "github.com/stas-tsarev/dz1_order_service/inventory/internal/repository/model"
	"github.com/stas-tsarev/dz1_order_service/inventory/internal/service/converter"
	"github.com/stas-tsarev/dz1_order_service/inventory/internal/service/model"
)

func (s *ServiceSuite) TestListParts_Success() {
	arrayLen := gofakeit.IntN(10)

	uuids := make([]string, arrayLen)
	names := make([]string, arrayLen)
	manufacturer := make([]string, arrayLen)
	tags := make([]string, arrayLen)

	for i := 0; i < arrayLen; i++ {
		uuids[i] = gofakeit.UUID()
		names[i] = gofakeit.Name()
		manufacturer[i] = gofakeit.Country()
		tags[i] = gofakeit.Name()
	}

	partFilter := &model.PartsFilter{
		Uuids:                 uuids,
		Names:                 names,
		Categories:            []model.Category{model.CATEGORY_PORTHOLE, model.CATEGORY_ENGINE},
		ManufacturerCountries: manufacturer,
		Tags:                  tags,
	}

	expectedParts := []model2.Part{
		{
			Uuid: gofakeit.UUID(),
			Info: model2.PartInfo{
				Name:          gofakeit.Name(),
				Description:   gofakeit.Word(),
				Price:         gofakeit.Price(1.00, 100000.00),
				StockQuantity: gofakeit.Int64(),
				Category:      model2.CATEGORY_PORTHOLE,
				Dimensions: model2.Dimensions{
					Length: gofakeit.Float64(),
					Width:  gofakeit.Float64(),
					Height: gofakeit.Float64(),
					Weight: gofakeit.Float64(),
				},
				Manufacturer: model2.Manufacturer{
					Name:    gofakeit.Name(),
					Country: gofakeit.Country(),
					WebSite: gofakeit.Name(),
				},
				Tags:     nil,
				Metadata: nil,
			},
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
		{
			Uuid: gofakeit.UUID(),
			Info: model2.PartInfo{
				Name:          gofakeit.Name(),
				Description:   gofakeit.Word(),
				Price:         gofakeit.Price(1.00, 100000.00),
				StockQuantity: gofakeit.Int64(),
				Category:      model2.CATEGORY_ENGINE,
				Dimensions: model2.Dimensions{
					Length: gofakeit.Float64(),
					Width:  gofakeit.Float64(),
					Height: gofakeit.Float64(),
					Weight: gofakeit.Float64(),
				},
				Manufacturer: model2.Manufacturer{
					Name:    gofakeit.Name(),
					Country: gofakeit.Country(),
					WebSite: gofakeit.Name(),
				},
				Tags:     nil,
				Metadata: nil,
			},
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	}

	s.InventoryRepository.On("ListParts", s.ctx, converter.ServiceFilterToRepoFilter(partFilter)).Return(expectedParts, nil)

	parts, err := s.service.ListParts(s.ctx, partFilter)

	newExpectedParts := make([]model.Part, len(expectedParts))
	for i := 0; i < len(expectedParts); i++ {
		newExpectedParts[i] = converter.RepoPartToServicePart(expectedParts[i])
	}

	s.Require().NoError(err)
	s.Require().Equal(newExpectedParts, parts)
}

func (s *ServiceSuite) TestListParts_NotFound() {
	partFilter := &model.PartsFilter{
		Uuids:                 []string{gofakeit.UUID(), gofakeit.UUID(), gofakeit.UUID()},
		Names:                 []string{gofakeit.Name(), gofakeit.Name(), gofakeit.Name()},
		Categories:            []model.Category{model.CATEGORY_FUEL, model.CATEGORY_WING},
		ManufacturerCountries: []string{gofakeit.Country(), gofakeit.Country()},
		Tags:                  []string{gofakeit.Word(), gofakeit.Word(), gofakeit.Word()},
	}

	expectedError := model2.ErrorPartNotFound

	s.InventoryRepository.On("ListParts", s.ctx, converter.ServiceFilterToRepoFilter(partFilter)).Return([]model2.Part{}, expectedError)

	parts, err := s.service.ListParts(s.ctx, partFilter)

	s.Require().Error(err)
	s.Require().Equal(err, expectedError)
	s.Require().Empty(parts)
}
