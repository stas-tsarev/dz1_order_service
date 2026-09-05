package inventory

import (
	"context"
	"log"

	"github.com/stas-tsarev/dz1_order_service/inventory/internal/service/converter"
	"github.com/stas-tsarev/dz1_order_service/inventory/internal/service/model"
)

func (s *service) CreatePart(ctx context.Context, partInfo model.PartInfo) (string, error) {
	newUuid, err := s.inventoryRepo.CreatePart(ctx, converter.ServicePartInfoToRepoPartInfo(partInfo))
	if err != nil {
		log.Printf("CreatePart error: %s", err)
		return "", err
	}

	log.Printf("CreatePart newUuid: %s", newUuid)
	return newUuid, nil
}
