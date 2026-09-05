package converter

import (
	repoModel "github.com/stas-tsarev/dz1_order_service/inventory/internal/repository/model"
	"github.com/stas-tsarev/dz1_order_service/inventory/internal/service/model"
)

/*
	repo -> service
*/

func RepoCategoryToServiceCategory(repoCategory repoModel.Category) model.Category {
	return model.Category(repoCategory)
}

func RepoDimensionsToServiceDimensions(repoDimensions repoModel.Dimensions) model.Dimensions {
	return model.Dimensions(repoDimensions)
}

func RepoManufacturerToServiceManufacturer(repoManufacturer repoModel.Manufacturer) model.Manufacturer {
	return model.Manufacturer(repoManufacturer)
}

func RepoValueToServiceValue(repoValue repoModel.Value) model.Value {
	return model.Value(repoValue)
}

func RepoMetadataToServiceMetadata(metadata map[string]repoModel.Value) map[string]model.Value {
	newMetadata := make(map[string]model.Value)
	for key, value := range metadata {
		newMetadata[key] = RepoValueToServiceValue(value)
	}

	return newMetadata
}

func RepoPartInfoToServicePartInfo(repoPartInfo repoModel.PartInfo) model.PartInfo {
	return model.PartInfo{
		Name:          repoPartInfo.Name,
		Description:   repoPartInfo.Description,
		Price:         repoPartInfo.Price,
		StockQuantity: repoPartInfo.StockQuantity,
		Category:      RepoCategoryToServiceCategory(repoPartInfo.Category),
		Dimensions:    RepoDimensionsToServiceDimensions(repoPartInfo.Dimensions),
		Manufacturer:  RepoManufacturerToServiceManufacturer(repoPartInfo.Manufacturer),
		Tags:          repoPartInfo.Tags,
		Metadata:      RepoMetadataToServiceMetadata(repoPartInfo.Metadata),
	}
}

func RepoPartToServicePart(repoPart repoModel.Part) model.Part {
	return model.Part{
		Uuid:      repoPart.Uuid,
		Info:      RepoPartInfoToServicePartInfo(repoPart.Info),
		CreatedAt: repoPart.CreatedAt,
		UpdatedAt: repoPart.UpdatedAt,
	}
}

/*
	service -> repo
*/

func ServiceCategoryToRepoCategory(serviceCategory model.Category) repoModel.Category {
	return repoModel.Category(serviceCategory)
}

func ServiceDimensionsToRepoDimensions(serviceDimensions model.Dimensions) repoModel.Dimensions {
	return repoModel.Dimensions(serviceDimensions)
}

func ServiceManufacturerToRepoManufacturer(serviceManufacturer model.Manufacturer) repoModel.Manufacturer {
	return repoModel.Manufacturer(serviceManufacturer)
}

func ServiceValueToRepoValue(serviceValue model.Value) repoModel.Value {
	return repoModel.Value(serviceValue)
}

func ServiceMetadataToRepoMetadata(metadata map[string]model.Value) map[string]repoModel.Value {
	newMetadata := make(map[string]repoModel.Value)
	for key, value := range metadata {
		newMetadata[key] = ServiceValueToRepoValue(value)
	}

	return newMetadata
}

func ServicePartInfoToRepoPartInfo(servicePart model.PartInfo) repoModel.PartInfo {
	return repoModel.PartInfo{
		Name:          servicePart.Name,
		Description:   servicePart.Description,
		Price:         servicePart.Price,
		StockQuantity: servicePart.StockQuantity,
		Category:      ServiceCategoryToRepoCategory(servicePart.Category),
		Dimensions:    ServiceDimensionsToRepoDimensions(servicePart.Dimensions),
		Manufacturer:  ServiceManufacturerToRepoManufacturer(servicePart.Manufacturer),
		Tags:          servicePart.Tags,
		Metadata:      ServiceMetadataToRepoMetadata(servicePart.Metadata),
	}
}

func ServiceFilterToRepoFilter(serviceFilter *model.PartsFilter) *repoModel.PartsFilter {
	newCategories := make([]repoModel.Category, 0)
	for _, category := range serviceFilter.Categories {
		newCategories = append(newCategories, ServiceCategoryToRepoCategory(category))
	}

	return &repoModel.PartsFilter{
		Uuids:                 serviceFilter.Uuids,
		Names:                 serviceFilter.Names,
		Categories:            newCategories,
		ManufacturerCountries: serviceFilter.ManufacturerCountries,
		Tags:                  serviceFilter.Tags,
	}
}
