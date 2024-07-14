package response

import (
	"daijai-service/models/dao"
	"github.com/shopspring/decimal"
)

type EstimateItemMaterial struct {
	MaterialAmount *decimal.Decimal `json:"materialAmount"`
	MaterialUnit   *string          `json:"materialUnit"`
	Material       []dao.Material   `json:"material"`
}

type EstimateItem struct {
	Id                   int              `json:"id"`
	Name                 string           `json:"name"`
	Code                 string           `json:"code"`
	Price                *decimal.Decimal `json:"price"`
	EstimateItemMaterial EstimateItemMaterial
}

type EstimateItemTypeResponse struct {
	Id           int          `json:"estimate_item_type_id"`
	Name         string       `json:"name"`
	EstimateItem EstimateItem `json:"estimateItem"`
}

type ProjectResponse struct {
	Id          int                      `json:"id"`
	ProjectName string                   `json:"projectName"`
	CreateBy    string                   `json:"createBy"`
	Detail      EstimateItemTypeResponse `json:"detail"`
}
