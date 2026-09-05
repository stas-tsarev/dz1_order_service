package v1

import (
	"context"

	"github.com/google/uuid"
	order_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/openapi/order/v1"
)

func (a *api) PayOrder(ctx context.Context, req order_v1.OptPayOrderRequest, params order_v1.PayOrderParams) (order_v1.PayOrderRes, error) {
	order, err := a.OrderService.GetOrder(ctx, params.OrderUUID.String())
	if err != nil {
		return &order_v1.R404NotFound{
			Code:    404,
			Message: err.Error(),
		}, nil
	}

	bytesMethod, err := req.Value.PaymentMethod.MarshalText()
	if err != nil {
		return nil, err
	}

	transactionUuid, err := a.OrderService.PayOrder(ctx, params.OrderUUID.String(), order.UserUuid, string(bytesMethod))
	if err != nil {
		return nil, err
	}

	newTransactionUuid, _ := uuid.Parse(transactionUuid)

	return &order_v1.PayOrderResponse{TransactionUUID: newTransactionUuid}, nil
}
