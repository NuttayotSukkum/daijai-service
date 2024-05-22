package dao

type MaterialFieldDetails struct {
	Id             uint `gorm:"primarykey"`
	MaterialFields string
	Name           string
	Code           string
}

func (MaterialFieldDetails) TableName() string {
	return "tbl_material_field_details"
}
