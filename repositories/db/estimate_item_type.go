package db

import (
	"daijai-service/models/dao"
	"daijai-service/repositories"
	"gorm.io/gorm"
)

type EstimateItemType struct {
	db *gorm.DB
}

func NewEstimateItemType(db *gorm.DB) repositories.EstimateItemType {
	return &EstimateItemType{db: db}
}

func (repo *EstimateItemType) Insert(itemType dao.EstimateItemTypes) error {
	if err := repo.db.Create(&itemType).Error; err != nil {
		return err
	}
	return nil
}

func (repo *EstimateItemType) FindEstimateItemTypesByName(name string) dao.EstimateItemTypes {
	if err := repo.db.Where("name = ?", name).First(&dao.EstimateItemTypes{}).Error; err != nil {
		return dao.EstimateItemTypes{}
	}
	return dao.EstimateItemTypes{}
}

func (repo *EstimateItemType) FindEstimateItemTypesById(id int) ([]dao.EstimateItemTypes, error) {
	var estimateItemTypes []dao.EstimateItemTypes
	if err := repo.db.Where("id = ?", id).Find(&estimateItemTypes).Error; err != nil {
		return nil, err
	}
	return estimateItemTypes, nil
}

func (repo *EstimateItemType) FindEstimateItemTypeAll() ([]dao.EstimateItemTypes, error) {
	var estimateItemTypeAll []dao.EstimateItemTypes
	if err := repo.db.Find(&estimateItemTypeAll).Error; err != nil {
		return nil, err
	}
	return estimateItemTypeAll, nil
}
