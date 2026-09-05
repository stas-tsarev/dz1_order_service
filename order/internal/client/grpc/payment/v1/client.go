package v1

import (
	def "github.com/stas-tsarev/dz1_order_service/order/internal/client/grpc"
	payment_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/proto/payment/v1"
)

var _ def.PaymentClient = (*client)(nil)

type client struct {
	generatedClient payment_v1.PaymentServiceClient
}

func NewClient(generatedClient payment_v1.PaymentServiceClient) *client {
	return &client{generatedClient: generatedClient}
}
