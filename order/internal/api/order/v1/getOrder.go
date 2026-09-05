package v1

import (
	"context"

	"github.com/google/uuid"
	order_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/openapi/order/v1"
)

func (a *api) GetOrder(ctx context.Context, params order_v1.GetOrderParams) (order_v1.GetOrderRes, error) {
	order, err := a.OrderService.GetOrder(ctx, params.OrderUUID.String())

	if err != nil {
		return &order_v1.R404NotFound{
			Code:    404,
			Message: err.Error(),
		}, nil
	}

	orderUuid, _ := uuid.Parse(order.OrderUuid)
	userUuid, _ := uuid.Parse(order.UserUuid)
	partsUuid := make([]uuid.UUID, 0, len(order.PartUuids))
	for _, partUuid := range order.PartUuids {
		newPartUuid, _ := uuid.Parse(partUuid)
		partsUuid = append(partsUuid, newPartUuid)
	}

	var transactionNilUuid order_v1.NilUUID
	if order.TransactionUuid != nil {
		transactionUuid, _ := uuid.Parse(*order.TransactionUuid)
		transactionNilUuid = order_v1.NilUUID{
			Value: transactionUuid,
			Null:  false,
		}
	} else {
		transactionNilUuid = order_v1.NilUUID{
			Null: true,
		}
	}

	var paymentMethod order_v1.PaymentMethod
	if order.PaymentMethod != nil {
		paymentMethod = order_v1.PaymentMethod(*order.PaymentMethod)
	} else {
		paymentMethod = order_v1.PaymentMethodUNKNOWN
	}

	return &order_v1.Order{
		OrderUUID:       orderUuid,
		UserUUID:        userUuid,
		PartUuids:       partsUuid,
		TotalPrice:      order.TotalPrice,
		TransactionUUID: transactionNilUuid,
		PaymentMethod:   paymentMethod,
		Status:          order_v1.OrderStatus(order.Status),
	}, nil
}
