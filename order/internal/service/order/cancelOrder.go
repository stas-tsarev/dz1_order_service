package order

import (
	"context"
	"errors"

	"github.com/stas-tsarev/dz1_order_service/order/internal/errs"
)

func (s *service) CancelOrder(ctx context.Context, orderUuid string) error {
	err := s.orderRepository.CancelOrder(ctx, orderUuid)

	if err != nil {
		if errors.Is(err, errs.ErrorOrderNotFound) {
			return errs.ErrorOrderNotFound
		}
		if errors.Is(err, errs.ErrorConflict) {
			return errs.ErrorConflict
		}
	}
	return nil
}
