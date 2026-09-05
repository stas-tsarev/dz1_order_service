package payment

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/stas-tsarev/dz1_order_service/payment/internal/model"
)

func (s *service) PayOrder(_ context.Context, _ model.PaymentInfo) (string, error) {
	uuid := uuid.New()
	log.Printf("Create payment with uuid: %s", uuid)

	return uuid.String(), nil
}
