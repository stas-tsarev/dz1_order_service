package service

import (
	"context"

	"github.com/stas-tsarev/dz1_order_service/order/internal/service/model"
)

type OrderService interface {
	CreateOrder(ctx context.Context, userUuid string, partUuids []string) (string, float64, error)
	PayOrder(ctx context.Context, orderUuid string, userUuid string, paymentMethod string) (string, error)
	GetOrder(ctx context.Context, orderUuid string) (model.Order, error)
	CancelOrder(ctx context.Context, orderUuid string) error
}
