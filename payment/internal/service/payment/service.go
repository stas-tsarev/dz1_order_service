package payment

import def "github.com/stas-tsarev/dz1_order_service/payment/internal/service"

var _ def.PaymentService = (*service)(nil)

type service struct{}

func NewService() *service {
	return &service{}
}
