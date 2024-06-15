package dao

import (
	"github.com/shopspring/decimal"
)

type EstimateItemMaterial struct {
	Id             int          `gorm:"primary_key"`
	EstimateItemId int          `gorm:"not null"`
	EstimateItem   EstimateItem `gorm:"foreignKey:EstimateItemId;references:Id"`
	MaterialId     int
	Material       Material         `gor,:"foreignKey:MaterialId;references:Id"`
	MaterialAmount *decimal.Decimal `gorm:"type:decimal(10,2)"`
	MaterialUnit   *string          `gorm:"size: 10"`
}

func (EstimateItemMaterial) TableName() string {
	return "tbl_estimate_item_material"
}
