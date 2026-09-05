package inventory

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/stas-tsarev/dz1_order_service/inventory/internal/repository/model"
)

func (r *repository) CreatePart(_ context.Context, partInfo model.PartInfo) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	newUUID := uuid.NewString()

	newPart := model.Part{
		Uuid:      newUUID,
		Info:      partInfo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	r.data[newUUID] = newPart

	return newUUID, nil
}
