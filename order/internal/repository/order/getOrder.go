package order

import (
	"context"

	"github.com/stas-tsarev/dz1_order_service/order/internal/errs"
	"github.com/stas-tsarev/dz1_order_service/order/internal/repository/model"
)

func (r *repository) GetOrder(_ context.Context, orderUuid string) (model.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	order, ok := r.data[orderUuid]
	if !ok {
		return model.Order{}, errs.ErrorOrderNotFound
	}

	return order, nil
}
