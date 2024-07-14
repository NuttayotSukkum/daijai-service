package utils

import (
	"daijai-service/models/dao"
	"daijai-service/models/response"
)

func ObjectMapper(object dao.EstimateItemMaterial, materials []dao.Material) response.ProjectResponse {

	estimateMaterialResponse := response.EstimateItemMaterial{
		MaterialAmount: object.MaterialAmount,
		MaterialUnit:   object.MaterialUnit,
		Material:       materials,
	}

	estimateItemResponse := response.EstimateItem{
		Id:                   object.EstimateItem.Id,
		Code:                 object.EstimateItem.Code,
		Name:                 object.EstimateItem.Name,
		Price:                object.EstimateItem.Price,
		EstimateItemMaterial: estimateMaterialResponse,
	}

	estimateItemMaterial := response.EstimateItemTypeResponse{
		Id:           object.EstimateItem.EstimateItemType.Id,
		Name:         object.EstimateItem.EstimateItemType.Name,
		EstimateItem: estimateItemResponse,
	}
	project := response.ProjectResponse{
		Id:          object.ProjectId,
		ProjectName: object.Project.ProjectName,
		CreateBy:    object.Project.CreatedBy,
		Detail:      estimateItemMaterial,
	}
	return project
}
