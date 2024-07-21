package dao

type Project struct {
	Id          int `gorm:"primary_key"`
	ProjectName string
	Status      string
	CreatedAt   string
	UpdatedAt   string
	CreatedBy   string
	Details     bool
}

func (Project) TableName() string {
	return "tbl_project"
}
