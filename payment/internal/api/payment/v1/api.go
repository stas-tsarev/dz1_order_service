package v1

import (
	"github.com/stas-tsarev/dz1_order_service/payment/internal/service"
	payment_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/proto/payment/v1"
)

type api struct {
	payment_v1.UnimplementedPaymentServiceServer

	PaymentService service.PaymentService
}

func NewApi(paymentService service.PaymentService) *api {
	return &api{
		PaymentService: paymentService,
	}
}
