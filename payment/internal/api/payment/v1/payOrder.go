package v1

import (
	"context"

	"github.com/stas-tsarev/dz1_order_service/payment/internal/converter"
	payment_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/proto/payment/v1"
)

func (a *api) PayOrder(ctx context.Context, req *payment_v1.PayOrderRequest) (*payment_v1.PayOrderResponse, error) {
	uuid, err := a.PaymentService.PayOrder(ctx, converter.PayOrderToPaymentInfo(req.GetPay()))
	if err != nil {
		return nil, err
	}

	return &payment_v1.PayOrderResponse{
		TransactionUuid: uuid,
	}, nil
}
