package dao

type Category3 struct {
	Id   int `gorm:"primary_key"`
	Name string
	Code string
}

func (Category3) TableName() string {
	return "tbl_category3s"
}
