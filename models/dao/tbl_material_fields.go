package dao

type MaterialFields struct {
	Id   uint `gorm:"primarykey"`
	Name string
}

func (MaterialFields) TableName() string {
	return "tbl_Material_Fields"
}
