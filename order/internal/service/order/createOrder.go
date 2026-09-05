package order

import (
	"context"

	model2 "github.com/stas-tsarev/dz1_order_service/order/internal/client/model"
	"github.com/stas-tsarev/dz1_order_service/order/internal/errs"
)

func (s *service) CreateOrder(ctx context.Context, userUuid string, partUuids []string) (string, float64, error) {
	parts, err := s.inventoryClient.ListParts(ctx, model2.PartsFilter{
		Uuids:                 partUuids,
		Names:                 nil,
		Categories:            nil,
		ManufacturerCountries: nil,
		Tags:                  nil,
	})

	if err != nil {
		return "", 0, errs.ErrorUnprocessableEntity
	}

	var totalPrice float64 = 0
	for _, part := range parts {
		totalPrice += part.Info.Price
	}

	orderUuid, _, _ := s.orderRepository.CreateOrder(ctx, userUuid, partUuids, totalPrice)

	return orderUuid, totalPrice, nil
}
