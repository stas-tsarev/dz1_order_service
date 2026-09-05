package v1

import (
	"context"
	"errors"

	"github.com/stas-tsarev/dz1_order_service/order/internal/errs"
	order_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/openapi/order/v1"
)

func (a *api) CancelOrder(ctx context.Context, params order_v1.CancelOrderParams) (order_v1.CancelOrderRes, error) {
	err := a.OrderService.CancelOrder(ctx, params.OrderUUID.String())

	if err != nil {
		if errors.Is(err, errs.ErrorOrderNotFound) {
			return &order_v1.R404NotFound{
				Code:    404,
				Message: err.Error(),
			}, nil
		}
		if errors.Is(err, errs.ErrorConflict) {
			return &order_v1.R409Conflict{
				Code:    409,
				Message: err.Error(),
			}, nil
		}
	}

	return &order_v1.CancelOrderNoContent{}, nil
}
