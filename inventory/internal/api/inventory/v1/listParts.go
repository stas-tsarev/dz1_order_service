package v1

import (
	"context"

	"github.com/stas-tsarev/dz1_order_service/inventory/internal/converter"
	inventory_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/proto/inventory/v1"
)

func (a *api) ListParts(ctx context.Context, req *inventory_v1.ListPartsRequest) (*inventory_v1.ListPartsResponse, error) {
	parts, err := a.InventoryService.ListParts(ctx, converter.ApiFilterToServiceFilter(req.Filter))
	if err != nil {
		return nil, err
	}

	result := make([]*inventory_v1.Part, 0, len(parts))
	for _, part := range parts {
		result = append(result, converter.ServicePartToApiPart(part))
	}

	return &inventory_v1.ListPartsResponse{
		Parts: result,
	}, nil
}
