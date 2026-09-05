package v1

import (
	"context"

	"github.com/stas-tsarev/dz1_order_service/order/internal/client/converter"
	"github.com/stas-tsarev/dz1_order_service/order/internal/client/model"
	inventory_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/proto/inventory/v1"
)

func (c *client) ListParts(ctx context.Context, filter model.PartsFilter) ([]model.Part, error) {
	parts, err := c.generatedClient.ListParts(ctx, &inventory_v1.ListPartsRequest{
		Filter: converter.ClientFilterToApiFilter(filter),
	})
	if err != nil {
		return nil, err
	}

	result := make([]model.Part, 0, len(parts.Parts))
	for _, part := range parts.Parts {
		result = append(result, converter.ApiPartToClientPart(part))
	}

	return result, nil
}
