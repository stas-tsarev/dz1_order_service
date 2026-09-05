package v1

import (
	"context"

	"github.com/google/uuid"
	order_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/openapi/order/v1"
)

func (a *api) CreateOrder(ctx context.Context, req order_v1.OptCreateOrderRequest) (order_v1.CreateOrderRes, error) {
	partsUuids := make([]string, 0, len(req.Value.PartUuids))
	for _, partUuid := range req.Value.PartUuids {
		partsUuids = append(partsUuids, partUuid.String())
	}
	orderUuid, totalPrice, err := a.OrderService.CreateOrder(ctx, req.Value.UserUUID.String(), partsUuids)

	if err != nil {
		return &order_v1.R422UnprocessableEntity{
			Code:    422,
			Message: err.Error(),
		}, nil
	}

	orderUuidResult, _ := uuid.Parse(orderUuid)
	return &order_v1.CreateOrderResponse{
		OrderUUID:  orderUuidResult,
		TotalPrice: totalPrice,
	}, nil
}
