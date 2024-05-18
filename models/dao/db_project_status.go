package dao

import "github.com/google/uuid"

type ProjectStatus struct {
	ID          uint `gorm:"primarykey"`
	ProjectId   uuid.UUID
	ProjectName string
	Status      string
	CreatedAt   string
	UpdatedAt   string
	CreatedBy   string
	Details     bool
}

func (ProjectStatus) TableName() string {
	return "tbl_project_status"
}
