package grpc

import (
	"context"

	"github.com/stas-tsarev/dz1_order_service/order/internal/client/model"
)

type InventoryClient interface {
	ListParts(ctx context.Context, filter model.PartsFilter) ([]model.Part, error)
}

type PaymentClient interface {
	PayOrder(ctx context.Context, orderUuid, userUuid, paymentMethod string) (string, error)
}
