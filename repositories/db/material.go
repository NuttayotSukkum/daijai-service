package db

import (
	"daijai-service/models/dao"
	"daijai-service/repositories"
	"fmt"
	"gorm.io/gorm"
)

type MaterialRepo struct {
	db *gorm.DB
}

func MaterialRepository(db *gorm.DB) repositories.Material {
	return &MaterialRepo{db: db}
}

func (repo *MaterialRepo) Insert(material dao.Material) (dao.Material, error) {
	if err := repo.db.Create(&material).Error; err != nil {
		return material, err
	}
	return material, nil
}

func (repo *MaterialRepo) Update(material dao.Material) (dao.Material, error) {
	if err := repo.db.Model(&material).Updates(&material).Error; err != nil {
		return material, err
	}
	var updatedMaterial dao.Material
	if err := repo.db.Where("id = ?", material.Id).Preload("Category3").First(&updatedMaterial).Error; err != nil {
		return updatedMaterial, err
	}

	return updatedMaterial, nil
}

func (repo *MaterialRepo) GetALL() ([]dao.Material, error) {
	var materials []dao.Material
	err := repo.db.Preload("Category3").Find(&materials).Error
	if err != nil {
		return materials, err
	}
	if len(materials) == 0 {
		return materials, fmt.Errorf("no material found")
	}
	return materials, nil
}
