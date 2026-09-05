package order

import (
	"context"

	"github.com/stas-tsarev/dz1_order_service/order/internal/errs"
)

func (s *service) PayOrder(ctx context.Context, orderUuid string, userUuid string, paymentMethod string) (string, error) {
	transactionUuid, err := s.paymentClient.PayOrder(ctx, orderUuid, userUuid, paymentMethod)
	if err != nil {
		return "", err
	}

	err = s.orderRepository.PayOrder(ctx, orderUuid, paymentMethod, transactionUuid)
	if err != nil {
		return "", errs.ErrorOrderNotFound
	}
	return transactionUuid, nil
}
