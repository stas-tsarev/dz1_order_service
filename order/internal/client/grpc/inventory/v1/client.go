package v1

import (
	def "github.com/stas-tsarev/dz1_order_service/order/internal/client/grpc"
	inventory_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/proto/inventory/v1"
)

var _ def.InventoryClient = (*client)(nil)

type client struct {
	generatedClient inventory_v1.InventoryServiceClient
}

func NewClient(generatedClient inventory_v1.InventoryServiceClient) *client {
	return &client{generatedClient: generatedClient}
}
