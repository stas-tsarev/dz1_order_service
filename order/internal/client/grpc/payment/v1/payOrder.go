package v1

import (
	"context"

	"github.com/stas-tsarev/dz1_order_service/order/internal/client/converter"
	payment_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/proto/payment/v1"
)

func (c *client) PayOrder(ctx context.Context, orderUuid, userUuid, paymentMethod string) (string, error) {
	result, err := c.generatedClient.PayOrder(ctx, &payment_v1.PayOrderRequest{
		Pay: &payment_v1.Pay{
			OrderUuid:     orderUuid,
			UserUuid:      userUuid,
			PaymentMethod: converter.ClientMethodToApiMethod(paymentMethod),
		},
	})

	if err != nil {
		return "", err
	}

	return result.TransactionUuid, nil
}
