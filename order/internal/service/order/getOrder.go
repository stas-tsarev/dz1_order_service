package order

import (
	"context"

	"github.com/stas-tsarev/dz1_order_service/order/internal/errs"
	"github.com/stas-tsarev/dz1_order_service/order/internal/service/converter"
	"github.com/stas-tsarev/dz1_order_service/order/internal/service/model"
)

func (s *service) GetOrder(ctx context.Context, orderUuid string) (model.Order, error) {
	order, err := s.orderRepository.GetOrder(ctx, orderUuid)
	if err != nil {
		return model.Order{}, errs.ErrorOrderNotFound
	}

	return converter.RepoOrderToServiceOrder(order), nil
}
