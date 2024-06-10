package dao

type Category3 struct {
	Id   int `gorm:"primarykey"`
	Name string
	Code string
}

func (Category3) TableName() string {
	return "tbl_category3s"
}
