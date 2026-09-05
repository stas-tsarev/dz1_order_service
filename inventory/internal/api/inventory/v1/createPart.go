package v1

import (
	"context"

	"github.com/stas-tsarev/dz1_order_service/inventory/internal/converter"
	inventory_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/proto/inventory/v1"
)

func (a *api) CreatePart(ctx context.Context, req *inventory_v1.CreatePartRequest) (*inventory_v1.CreatePartResponse, error) {
	newUuid, err := a.InventoryService.CreatePart(ctx, converter.ApiPartInfoToServicePartInfo(req.Info))
	if err != nil {
		return nil, err
	}

	return &inventory_v1.CreatePartResponse{Uuid: newUuid}, nil
}
