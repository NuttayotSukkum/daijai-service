package requests

import "github.com/shopspring/decimal"

type RequestEstimateItemMaterial struct {
	EstimateItemId int              `json:"estimate_item_id"`
	MaterialId     []int            `json:"material_id"`
	MaterialAmount *decimal.Decimal `json:"material_amount"`
	MaterialUnit   *string          `json:"material_unit"`
}
