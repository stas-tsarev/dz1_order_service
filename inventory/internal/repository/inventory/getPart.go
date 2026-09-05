package inventory

import (
	"context"

	"github.com/stas-tsarev/dz1_order_service/inventory/internal/repository/model"
)

func (r *repository) GetPart(_ context.Context, uuid string) (model.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	repoPart, ok := r.data[uuid]
	if !ok {
		return model.Part{}, model.ErrorPartNotFound
	}

	return repoPart, nil
}
