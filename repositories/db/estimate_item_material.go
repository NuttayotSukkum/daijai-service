package db

import (
	"daijai-service/models/dao"
	"daijai-service/repositories"
	"gorm.io/gorm"
	"log"
)

type EstimateItemMaterial struct {
	db *gorm.DB
}

func NewEstimateItemMaterial(db *gorm.DB) repositories.EstimateItemMaterial {
	return &EstimateItemMaterial{db: db}
}

func (repo EstimateItemMaterial) Insert(material dao.EstimateItemMaterial) (dao.EstimateItemMaterial, error) {
	if err := repo.db.Create(&material).Error; err != nil {
		return dao.EstimateItemMaterial{}, err
	}
	var response dao.EstimateItemMaterial
	if err := repo.db.Where("id = ?", material.Id).Preload("EstimateItem").Preload("EstimateItem.EstimateItemType").First(&response).Error; err != nil {
		return dao.EstimateItemMaterial{}, err
	}
	log.Println(response)
	return response, nil

}

//func (repo *EstimateItemType) GetAll() (dao.EstimateItemTypes, error) {
//
//}
