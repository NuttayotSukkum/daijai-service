package dao

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	ProjectName string
	CreatedBy   string
}
