package order

import (
	"context"

	"github.com/stas-tsarev/dz1_order_service/order/internal/errs"
	"github.com/stas-tsarev/dz1_order_service/order/internal/repository/model"
)

func (r *repository) PayOrder(_ context.Context, orderUuid string, paymentMethod string, transactionUuid string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	tmp, ok := r.data[orderUuid]

	if !ok {
		return errs.ErrorOrderNotFound
	}

	r.data[orderUuid] = model.Order{
		OrderUuid:       orderUuid,
		UserUuid:        tmp.UserUuid,
		PartUuids:       tmp.PartUuids,
		TotalPrice:      tmp.TotalPrice,
		TransactionUuid: &transactionUuid,
		PaymentMethod:   &paymentMethod,
		Status:          "PAID",
	}

	return nil
}
