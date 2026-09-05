package service

import (
	"context"

	"github.com/stas-tsarev/dz1_order_service/payment/internal/model"
)

type PaymentService interface {
	PayOrder(ctx context.Context, info model.PaymentInfo) (string, error)
}
