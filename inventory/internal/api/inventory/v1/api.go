package v1

import (
	"github.com/stas-tsarev/dz1_order_service/inventory/internal/service"
	inventory_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/proto/inventory/v1"
)

type api struct {
	inventory_v1.UnimplementedInventoryServiceServer
	InventoryService service.InventoryService
}

func NewApi(InventoryService service.InventoryService) *api {
	return &api{
		InventoryService: InventoryService,
	}
}
