package inventory

import (
	"context"

	"github.com/stas-tsarev/dz1_order_service/inventory/internal/repository/model"
)

func (r *repository) ListParts(_ context.Context, filter *model.PartsFilter) ([]model.Part, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	parts := make([]model.Part, 0, len(r.data))
	for _, part := range r.data {
		parts = append(parts, part)
	}

	if filter == nil {
		return parts, nil
	}

	if filter.Uuids == nil &&
		filter.Names == nil &&
		len(filter.Categories) == 0 &&
		filter.ManufacturerCountries == nil &&
		filter.Tags == nil {
		return parts, nil
	}

	if filter.Uuids != nil && len(parts) != 0 {
		uuidSet := make(map[string]bool, len(filter.Uuids))
		for _, uuid := range filter.Uuids {
			uuidSet[uuid] = true
		}
		result := make([]model.Part, 0, len(parts))
		for _, part := range parts {
			if uuidSet[part.Uuid] {
				result = append(result, part)
			}
		}
		parts = result
	}

	if filter.Names != nil && len(parts) != 0 {
		nameSet := make(map[string]bool, len(filter.Names))
		for _, name := range filter.Names {
			nameSet[name] = true
		}
		result := make([]model.Part, 0, len(parts))
		for _, part := range parts {
			if nameSet[part.Info.Name] {
				result = append(result, part)
			}
		}
		parts = result
	}

	if len(filter.Categories) != 0 && len(parts) != 0 {
		categorySet := make(map[string]bool, len(filter.Categories))
		for _, category := range filter.Categories {
			categorySet[category.String()] = true
		}
		result := make([]model.Part, 0, len(parts))
		for _, part := range parts {
			if categorySet[part.Info.Category.String()] {
				result = append(result, part)
			}
		}
		parts = result
	}

	if filter.ManufacturerCountries != nil && len(parts) != 0 {
		countrySet := make(map[string]bool, len(filter.ManufacturerCountries))
		for _, country := range filter.ManufacturerCountries {
			countrySet[country] = true
		}
		result := make([]model.Part, 0, len(parts))
		for _, part := range parts {
			if countrySet[part.Info.Manufacturer.Country] {
				result = append(result, part)
			}
		}
		parts = result
	}

	if filter.Tags != nil && len(parts) != 0 {
		tagSet := make(map[string]bool, len(filter.Tags))
		for _, tag := range filter.Tags {
			tagSet[tag] = true
		}
		result := make([]model.Part, 0, len(parts))
		for _, part := range parts {
			for ind := range part.Info.Tags {
				if tagSet[part.Info.Tags[ind]] {
					result = append(result, part)
					break
				}
			}
		}
		parts = result
	}

	if len(parts) == 0 {
		return []model.Part{}, model.ErrorPartNotFound
	}

	return parts, nil
}
