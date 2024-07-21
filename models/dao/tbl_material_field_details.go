package dao

type MaterialFieldDetail struct {
	Id              int `gorm:"primaryKey"`
	MaterialFieldId int
	MaterialField   MaterialField `gorm:"foreignKey:MaterialFieldId;references:Id"`
	Name            string
	Code            string
}

func (MaterialFieldDetail) TableName() string {
	return "tbl_material_field_details"
}
