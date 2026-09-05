package converter

import (
	"github.com/stas-tsarev/dz1_order_service/inventory/internal/service/model"
	inventory_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/proto/inventory/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

/*
	service -> api
*/

func ServiceManufacturerToApiManufacturer(manufacturer model.Manufacturer) *inventory_v1.Manufacturer {
	return &inventory_v1.Manufacturer{
		Name:    manufacturer.Name,
		Country: manufacturer.Country,
		Website: manufacturer.WebSite,
	}
}

func ServiceDimensionsToApiDimensions(serviceDimensions model.Dimensions) *inventory_v1.Dimensions {
	return &inventory_v1.Dimensions{
		Length: serviceDimensions.Length,
		Width:  serviceDimensions.Width,
		Height: serviceDimensions.Height,
		Weight: serviceDimensions.Weight,
	}
}

func ServiceCategoryToApiCategory(serviceCategory model.Category) inventory_v1.Category {
	return inventory_v1.Category(serviceCategory)
}

func ServiceValueToApiValue(serviceValue model.Value) *inventory_v1.Value {
	value := &inventory_v1.Value{}

	switch v := serviceValue.(type) {
	case model.StringValue:
		value.StringValue = v.Value
	case *model.StringValue:
		value.StringValue = v.Value
	case model.Int64Value:
		value.Int64Value = v.Value
	case *model.Int64Value:
		value.Int64Value = v.Value
	case model.DoubleValue:
		value.DoubleValue = v.Value
	case *model.DoubleValue:
		value.DoubleValue = v.Value
	case model.BoolValue:
		value.BoolValue = v.Value
	case *model.BoolValue:
		value.BoolValue = v.Value
	}

	return value
}

func ServiceMetadataToApiMetadata(metadata map[string]model.Value) map[string]*inventory_v1.Value {
	newMetadata := make(map[string]*inventory_v1.Value)
	for key, value := range metadata {
		newMetadata[key] = ServiceValueToApiValue(value)
	}
	return newMetadata
}

func ServicePartInfoToApiPartInfo(partInfo model.PartInfo) *inventory_v1.PartInfo {
	return &inventory_v1.PartInfo{
		Name:          partInfo.Name,
		Description:   partInfo.Description,
		Price:         partInfo.Price,
		StockQuantity: partInfo.StockQuantity,
		Category:      ServiceCategoryToApiCategory(partInfo.Category),
		Dimensions:    ServiceDimensionsToApiDimensions(partInfo.Dimensions),
		Manufacturer:  ServiceManufacturerToApiManufacturer(partInfo.Manufacturer),
		Tags:          partInfo.Tags,
		Metadata:      ServiceMetadataToApiMetadata(partInfo.Metadata),
	}
}

func ServicePartToApiPart(servicePart model.Part) *inventory_v1.Part {
	return &inventory_v1.Part{
		Uuid:      servicePart.Uuid,
		Info:      ServicePartInfoToApiPartInfo(servicePart.Info),
		CreatedAt: timestamppb.New(servicePart.CreatedAt),
		UpdatedAt: timestamppb.New(servicePart.UpdatedAt),
	}
}

/*
	api -> service
*/

func ApiManufacturerToServiceManufacturer(manufacturer *inventory_v1.Manufacturer) model.Manufacturer {
	return model.Manufacturer{
		Name:    manufacturer.Name,
		Country: manufacturer.Country,
		WebSite: manufacturer.Website,
	}
}

func ApiDimensionsToServiceDimensions(apiDimensions *inventory_v1.Dimensions) model.Dimensions {
	return model.Dimensions{
		Length: apiDimensions.Length,
		Width:  apiDimensions.Width,
		Height: apiDimensions.Height,
		Weight: apiDimensions.Weight,
	}
}

func ApiCategoryToServiceCategory(category inventory_v1.Category) model.Category {
	return model.Category(category)
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

func ApiPartInfoToServicePartInfo(partInfo *inventory_v1.PartInfo) model.PartInfo {
	return model.PartInfo{
		Name:          partInfo.Name,
		Description:   partInfo.Description,
		Price:         partInfo.Price,
		StockQuantity: partInfo.StockQuantity,
		Category:      ApiCategoryToServiceCategory(partInfo.Category),
		Dimensions:    ApiDimensionsToServiceDimensions(partInfo.Dimensions),
		Manufacturer:  ApiManufacturerToServiceManufacturer(partInfo.Manufacturer),
		Tags:          partInfo.Tags,
		Metadata:      ApiMetadataToServiceMetadata(partInfo.Metadata),
	}
}

func ApiFilterToServiceFilter(filter *inventory_v1.PartsFilter) *model.PartsFilter {
	newCategories := make([]model.Category, 0)
	for _, category := range filter.Categories {
		newCategories = append(newCategories, ApiCategoryToServiceCategory(category))
	}

	return &model.PartsFilter{
		Uuids:                 filter.Uuids,
		Names:                 filter.Names,
		Categories:            newCategories,
		ManufacturerCountries: filter.ManufacturerCountries,
		Tags:                  filter.Tags,
	}
}
