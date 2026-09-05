package repository

import (
	"context"

	"github.com/stas-tsarev/dz1_order_service/order/internal/repository/model"
)

type OrderRepository interface {
	CreateOrder(_ context.Context, userUuid string, partUuids []string, totalPrice float64) (string, float64, error)
	PayOrder(ctx context.Context, orderUuid string, paymentMethod string, transactionUuid string) error
	GetOrder(ctx context.Context, orderUuid string) (model.Order, error)
	CancelOrder(ctx context.Context, orderUuid string) error
}
