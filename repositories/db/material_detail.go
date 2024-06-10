package db

import (
	"daijai-service/models/dao"
	"gorm.io/gorm"
	"log"
)

type MaterialDetailsRepo struct {
	db *gorm.DB
}

func NewMaterialDetails(db *gorm.DB) *MaterialDetailsRepo {
	return &MaterialDetailsRepo{db: db}
}

func (repo MaterialDetailsRepo) Insert(materialDetail dao.MaterialDetail) (dao.MaterialDetail, error) {
	if err := repo.db.Create(&materialDetail).Error; err != nil {
		return materialDetail, err
	}
	log.Println("Successfully inserted material detail")
	var response dao.MaterialDetail
	if err := repo.db.Where("id = ?", materialDetail.Id).Preload("Material").First(&response).Error; err != nil {
		return dao.MaterialDetail{}, err
	}
	return response, nil

}
