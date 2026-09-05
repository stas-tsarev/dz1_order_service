package order

import (
	"context"

	"github.com/stas-tsarev/dz1_order_service/order/internal/errs"
)

func (r *repository) CancelOrder(_ context.Context, orderUuid string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	order, ok := r.data[orderUuid]
	if !ok {
		return errs.ErrorOrderNotFound
	}

	if order.Status == "PAID" {
		return errs.ErrorConflict
	}

	order.Status = "CANCELED"
	r.data[orderUuid] = order

	return nil
}
