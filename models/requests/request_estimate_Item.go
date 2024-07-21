package requests

import "github.com/shopspring/decimal"

type RequestEstimateItem struct {
	EstimateItemTypeId int              `json:"estimate_item_type_id"`
	Name               *string          `json:"name"`
	Code               *string          `json:"code"`
	Price              *decimal.Decimal `json:"price"`
}
