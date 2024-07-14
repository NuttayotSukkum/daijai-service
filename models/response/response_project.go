package response

import (
	"github.com/shopspring/decimal"
)

type EstimateItemAll struct {
	Id                   int                    `json:"id"`
	Name                 string                 `json:"name"`
	Code                 string                 `json:"code"`
	Price                *decimal.Decimal       `json:"price"`
	EstimateItemMaterial []EstimateItemMaterial `json:"estimateItemMaterial"`
}

type EstimateItemTypeResponseAll struct {
	Id           int               `json:"estimate_item_type_id"`
	Name         string            `json:"name"`
	EstimateItem []EstimateItemAll `json:"estimateItem"`
}

type ProjectResponseList struct {
	Id          int                           `json:"id"`
	ProjectName string                        `json:"projectName"`
	CreateBy    string                        `json:"createBy"`
	Detail      []EstimateItemTypeResponseAll `json:"detail"`
}
