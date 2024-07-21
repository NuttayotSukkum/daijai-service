package db

import (
	"daijai-service/models/dao"
	"daijai-service/repositories"
	"gorm.io/gorm"
)

type ProjectEstimateItem struct {
	db *gorm.DB
}

func NewProjectEstimateItemRepository(db *gorm.DB) repositories.ProjectEstimateItem {
	return &ProjectEstimateItem{db: db}
}

func (repo *ProjectEstimateItem) Insert(req dao.ProjectEstimateItem) error {
	err := repo.db.Create(&req).Error
	if err != nil {
		return err
	}
	return nil
}
