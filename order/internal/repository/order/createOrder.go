package order

import (
	"context"

	"github.com/google/uuid"
	"github.com/stas-tsarev/dz1_order_service/order/internal/repository/model"
)

func (r *repository) CreateOrder(_ context.Context, userUuid string, partUuids []string, totalPrice float64) (string, float64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	newUuid := uuid.NewString()

	r.data[newUuid] = model.Order{
		OrderUuid:       newUuid,
		UserUuid:        userUuid,
		PartUuids:       partUuids,
		TotalPrice:      totalPrice,
		TransactionUuid: nil,
		PaymentMethod:   nil,
		Status:          "PENDING_PAYMENT",
	}

	return newUuid, totalPrice, nil
}
