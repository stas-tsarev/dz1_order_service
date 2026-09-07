package payment

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/stas-tsarev/dz1_order_service/payment/internal/model"
)

func (s *ServiceSuite) TestPayOrder_Success() {
	var (
		newOrderUuid = uuid.New().String()
		newUserUuid  = uuid.New().String()

		methods = []model.Method{
			model.PAYMENT_METHOD_UNKNOWN_UNSPECIFIED,
			model.PAYMENT_METHOD_CARD,
			model.PAYMENT_METHOD_SBP,
			model.PAYMENT_METHOD_CREDIT_CARD,
			model.PAYMENT_METHOD_INVESTOR_MONEY,
		}
		randMethod = gofakeit.Number(0, len(methods)-1)

		paymentInfo = model.PaymentInfo{
			OrderUuid:     newOrderUuid,
			UserUuid:      newUserUuid,
			PaymentMethod: methods[randMethod],
		}
	)

	res, err := s.service.PayOrder(s.ctx, paymentInfo)
	s.NoError(err)
	s.NotEmpty(res)
	s.IsType("string", res)
	s.Len(res, 36)
}
