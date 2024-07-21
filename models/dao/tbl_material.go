package dao

type Material struct {
	Id          int `gorm:"primarykey"`
	Category3Id int
	Category3   Category3 `gorm:"foreignKey:Category3Id;references:Id"`
	Code        *string
	Description *string
}

func (Material) TableName() string { return "tbl_materials" }
