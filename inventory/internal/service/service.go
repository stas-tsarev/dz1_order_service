package service

import (
	"context"

	"github.com/stas-tsarev/dz1_order_service/inventory/internal/service/model"
)

type InventoryService interface {
	GetPart(ctx context.Context, uuid string) (model.Part, error)
	ListParts(ctx context.Context, filter *model.PartsFilter) ([]model.Part, error)
	CreatePart(ctx context.Context, part model.PartInfo) (string, error)
}
