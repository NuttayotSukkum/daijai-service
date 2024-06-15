package dao

import "github.com/shopspring/decimal"

type EstimateItem struct {
	Id                 int `gorm:"primary_key;AUTO_INCREMENT"`
	EstimateItemTypeId int
	EstimateItemType   EstimateItemTypes `gorm:"foreign_key:estimate_item_type_id,references:Id;"`
	Name               string            `gorm:"size:50;"`
	Code               string            `gorm:"size:50;"`
	Price              *decimal.Decimal  `gorm:"type:decimal(10,2);"`
}

func (EstimateItem) TableName() string {
	return "tbl_estimate_item"
}
