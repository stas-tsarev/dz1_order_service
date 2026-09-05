package inventory

import (
	"sync"

	def "github.com/stas-tsarev/dz1_order_service/inventory/internal/repository"
	"github.com/stas-tsarev/dz1_order_service/inventory/internal/repository/model"
)

var _ def.InventoryRepository = (*repository)(nil)

type repository struct {
	mu   sync.RWMutex
	data map[string]model.Part
}

func NewRepository() *repository {
	return &repository{
		data: make(map[string]model.Part),
	}
}
