package dao

type ProjectEstimateItem struct {
	Id             int64 `gorm:"primaryKey"`
	ProjectId      int64
	Project        Project `gorm:"foreignKey:ProjectId;references:Id"`
	EstimateItemId int
	EstimateItem   EstimateItem `gorm:"foreignKey:EstimateItemId;references:Id"`
}

func (ProjectEstimateItem) TableName() string {
	return "tb_project_estimate_item"
}
