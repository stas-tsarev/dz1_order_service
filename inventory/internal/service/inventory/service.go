package inventory

import (
	"github.com/stas-tsarev/dz1_order_service/inventory/internal/repository"
	def "github.com/stas-tsarev/dz1_order_service/inventory/internal/service"
)

var _ def.InventoryService = (*service)(nil)

type service struct {
	inventoryRepo repository.InventoryRepository
}

func NewService(inventoryRepo repository.InventoryRepository) *service {
	return &service{
		inventoryRepo: inventoryRepo,
	}
}
