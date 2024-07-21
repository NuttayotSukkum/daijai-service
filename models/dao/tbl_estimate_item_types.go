package dao

type EstimateItemTypes struct {
	Id   int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string
}

func (EstimateItemTypes) TableName() string {
	return "tbl_estimate_item_types"
}
