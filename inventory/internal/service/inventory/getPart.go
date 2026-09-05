package inventory

import (
	"context"
	"log"

	"github.com/stas-tsarev/dz1_order_service/inventory/internal/service/converter"
	"github.com/stas-tsarev/dz1_order_service/inventory/internal/service/model"
)

func (s *service) GetPart(ctx context.Context, id string) (model.Part, error) {
	part, err := s.inventoryRepo.GetPart(ctx, id)
	if err != nil {
		log.Printf("GetPart error: %s", err)
		return model.Part{}, err
	}

	log.Printf("GetPart: %v", part.Uuid)
	return converter.RepoPartToServicePart(part), nil
}
