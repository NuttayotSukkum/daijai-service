package db

import (
	"daijai-service/models/dao"
	"daijai-service/repositories"
	"gorm.io/gorm"
)

type EstimateItem struct {
	db *gorm.DB
}

func NewEstimateItem(db *gorm.DB) repositories.EstimateItem {
	return &EstimateItem{db: db}
}

func (repo *EstimateItem) Insert(estimateItem dao.EstimateItem) error {
	return repo.db.Create(&estimateItem).Error
}

func (repo *EstimateItem) FindEstimateItemAll() ([]dao.EstimateItem, error) {
	var estimateItems []dao.EstimateItem
	err := repo.db.Preload("EstimateItemType").Find(&estimateItems).Error
	if err != nil {
		return nil, err
	}
	return estimateItems, nil
}

func (repo *EstimateItem) FindEstimateItemExist(id int) []dao.EstimateItem {
	if err := repo.db.Where("id = ?", id).Find(&dao.EstimateItem{}).Error; err != nil {
		return nil
	}
	return []dao.EstimateItem{}
}
