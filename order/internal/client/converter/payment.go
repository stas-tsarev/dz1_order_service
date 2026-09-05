package converter

import (
	payment_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/proto/payment/v1"
)

func ClientMethodToApiMethod(method string) payment_v1.PaymentMethod {
	switch method {
	case "PAYMENT_METHOD_UNKNOWN_UNSPECIFIED":
		return payment_v1.PaymentMethod_PAYMENT_METHOD_UNKNOWN_UNSPECIFIED // Неизвестный способ
	case "PAYMENT_METHOD_CARD":
		return payment_v1.PaymentMethod_PAYMENT_METHOD_CARD // Банковская карта
	case "PAYMENT_METHOD_SBP":
		return payment_v1.PaymentMethod_PAYMENT_METHOD_SBP // Система быстрых платежей
	case "PAYMENT_METHOD_CREDIT_CARD":
		return payment_v1.PaymentMethod_PAYMENT_METHOD_CREDIT_CARD // Кредитная карта
	case "PAYMENT_METHOD_INVESTOR_MONEY":
		return payment_v1.PaymentMethod_PAYMENT_METHOD_INVESTOR_MONEY // Деньги инвестора (внутренний метод)
	}
	return payment_v1.PaymentMethod_PAYMENT_METHOD_UNKNOWN_UNSPECIFIED
}
