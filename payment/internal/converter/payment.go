package converter

import (
	"github.com/stas-tsarev/dz1_order_service/payment/internal/model"
	payment_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/proto/payment/v1"
)

func PayOrderToPaymentInfo(order *payment_v1.Pay) model.PaymentInfo {
	return model.PaymentInfo{
		OrderUuid:     order.OrderUuid,
		UserUuid:      order.UserUuid,
		PaymentMethod: model.Method(order.PaymentMethod),
	}
}
