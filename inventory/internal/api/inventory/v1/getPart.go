package v1

import (
	"context"

	"github.com/stas-tsarev/dz1_order_service/inventory/internal/converter"
	inventory_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/proto/inventory/v1"
)

func (a *api) GetPart(ctx context.Context, req *inventory_v1.GetPartRequest) (*inventory_v1.GetPartResponse, error) {
	part, err := a.InventoryService.GetPart(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}

	return &inventory_v1.GetPartResponse{
		Part: converter.ServicePartToApiPart(part),
	}, nil
}
