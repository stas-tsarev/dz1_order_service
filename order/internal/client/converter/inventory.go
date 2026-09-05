package converter

import (
	"github.com/stas-tsarev/dz1_order_service/order/internal/client/model"
	inventory_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/proto/inventory/v1"
)

/*
	client -> api
*/

func ClientCategoryToApiCategory(category model.Category) inventory_v1.Category {
	return inventory_v1.Category(category)
}

func ClientFilterToApiFilter(filter model.PartsFilter) *inventory_v1.PartsFilter {
	newCategories := make([]inventory_v1.Category, 0, len(filter.Categories))
	for _, category := range filter.Categories {
		newCategories = append(newCategories, ClientCategoryToApiCategory(category))
	}

	return &inventory_v1.PartsFilter{
		Uuids:                 filter.Uuids,
		Names:                 filter.Names,
		Categories:            newCategories,
		ManufacturerCountries: filter.ManufacturerCountries,
		Tags:                  filter.Tags,
	}
}

/*
	api -> client
*/

func ApiCategoryToClientCategory(category inventory_v1.Category) model.Category {
	return model.Category(category)
}

func ApiDimensionsToClientDimensions(dimension *inventory_v1.Dimensions) model.Dimensions {
	return model.Dimensions{
		Length: dimension.Length,
		Width:  dimension.Width,
		Height: dimension.Height,
		Weight: dimension.Weight,
	}
}

func ApiManufacturerToClientManufacturer(manufacturer *inventory_v1.Manufacturer) model.Manufacturer {
	return model.Manufacturer{
		Name:    manufacturer.Name,
		Country: manufacturer.Country,
		WebSite: manufacturer.Website,
	}
}

func ApiValueToServiceValue(apiValue *inventory_v1.Value) model.Value {
	// Проверяем, какое поле заполнено
	if apiValue.StringValue != "" {
		return model.StringValue{Value: apiValue.StringValue}
	}
	if apiValue.Int64Value != 0 {
		return model.Int64Value{Value: apiValue.Int64Value}
	}
	if apiValue.DoubleValue != 0 {
		return model.DoubleValue{Value: apiValue.DoubleValue}
	}
	if apiValue.BoolValue {
		return model.BoolValue{Value: apiValue.BoolValue}
	}

	// Если ничего не заполнено, возвращаем пустую строку
	return model.StringValue{Value: ""}
}

func ApiMetadataToServiceMetadata(metadata map[string]*inventory_v1.Value) map[string]model.Value {
	newMetadata := make(map[string]model.Value)
	for key, value := range metadata {
		newMetadata[key] = ApiValueToServiceValue(value)
	}
	return newMetadata
}

func ApiPartInfoToClientPartInfo(partInfo *inventory_v1.PartInfo) model.PartInfo {
	return model.PartInfo{
		Name:          partInfo.Name,
		Description:   partInfo.Description,
		Price:         partInfo.Price,
		StockQuantity: partInfo.StockQuantity,
		Category:      ApiCategoryToClientCategory(partInfo.Category),
		Dimensions:    ApiDimensionsToClientDimensions(partInfo.Dimensions),
		Manufacturer:  ApiManufacturerToClientManufacturer(partInfo.Manufacturer),
		Tags:          partInfo.Tags,
		Metadata:      ApiMetadataToServiceMetadata(partInfo.Metadata),
	}
}

func ApiPartToClientPart(part *inventory_v1.Part) model.Part {
	return model.Part{
		Uuid:      part.Uuid,
		Info:      ApiPartInfoToClientPartInfo(part.Info),
		CreatedAt: part.CreatedAt.AsTime(),
		UpdatedAt: part.UpdatedAt.AsTime(),
	}
}
