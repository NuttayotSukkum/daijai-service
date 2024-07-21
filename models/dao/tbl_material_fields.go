package dao

type MaterialField struct {
	Id   int `gorm:"primarykey"`
	Name string
}

func (MaterialField) TableName() string {
	return "tbl_Material_Fields"
}
