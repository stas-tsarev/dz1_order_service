package model

type Method int

const (
	PAYMENT_METHOD_UNKNOWN_UNSPECIFIED Method = iota // Неизвестный способ
	PAYMENT_METHOD_CARD                              // Банковская карта
	PAYMENT_METHOD_SBP                               // Система быстрых платежей
	PAYMENT_METHOD_CREDIT_CARD                       // Кредитная карта
	PAYMENT_METHOD_INVESTOR_MONEY                    // Деньги инвестора (внутренний метод)
)

func (i Method) String() string {
	switch i {
	case PAYMENT_METHOD_UNKNOWN_UNSPECIFIED:
		return "PAYMENT_METHOD_UNKNOWN_UNSPECIFIED"
	case PAYMENT_METHOD_CARD:
		return "PAYMENT_METHOD_CARD"
	case PAYMENT_METHOD_SBP:
		return "PAYMENT_METHOD_SBP"
	case PAYMENT_METHOD_CREDIT_CARD:
		return "PAYMENT_METHOD_CREDIT_CARD"
	case PAYMENT_METHOD_INVESTOR_MONEY:
		return "PAYMENT_METHOD_INVESTOR_MONEY"
	}
	return "PAYMENT_METHOD_UNKNOWN_UNSPECIFIED"
}

type PaymentInfo struct {
	OrderUuid     string
	UserUuid      string
	PaymentMethod Method
}
