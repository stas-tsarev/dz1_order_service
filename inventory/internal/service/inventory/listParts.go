package inventory

import (
	"context"
	"log"

	"github.com/stas-tsarev/dz1_order_service/inventory/internal/service/converter"
	"github.com/stas-tsarev/dz1_order_service/inventory/internal/service/model"
)

func (s *service) ListParts(ctx context.Context, filter *model.PartsFilter) ([]model.Part, error) {
	parts, err := s.inventoryRepo.ListParts(ctx, converter.ServiceFilterToRepoFilter(filter))
	if err != nil {
		log.Printf("ListParts error: %s", err)
		return []model.Part{}, err
	}

	log.Printf("ListParts completed")
	result := make([]model.Part, 0, len(parts))
	for _, part := range parts {
		result = append(result, converter.RepoPartToServicePart(part))
	}

	return result, nil
}
