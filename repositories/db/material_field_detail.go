package db

import (
	"daijai-service/models/dao"
	"daijai-service/repositories"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
)

type MaterialDetailRepo struct {
	db *gorm.DB
}

func NewMaterialDetailsRepository(db *gorm.DB) repositories.MaterialFieldsDetail {
	return &MaterialDetailRepo{db: db}
}

func (repo MaterialDetailRepo) Insert(req dao.MaterialFieldDetails) error {
	err := repo.db.Create(&req).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo MaterialDetailRepo) CheckExisting(name string) bool {
	var response dao.MaterialFieldDetails
	err := repo.db.Where("name = ?", name).First(&response).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
		return false
	}
	return true
}

func (repo MaterialDetailRepo) GetMaterialAll() []dao.MaterialFieldDetails {
	var response []dao.MaterialFieldDetails
	err := repo.db.Find(&response).Error
	if err != nil {
		return nil
	}
	return response
}

func (repo MaterialDetailRepo) GetMaterialById(id uint) dao.MaterialFieldDetails {
	var response dao.MaterialFieldDetails
	err := repo.db.Where("id = ?", id).First(&response).Error
	if err != nil {
		log.Println("error:{}", err)
		return dao.MaterialFieldDetails{}
	}
	return response
}

func (repo MaterialDetailRepo) GetMaterialByMaterialFields(id string) dao.MaterialFieldDetails {
	var response dao.MaterialFieldDetails
	err := repo.db.Where("material_fields = ?", id).First(&response).Error
	if err != nil {
		log.Println("error:{}", err)
		return dao.MaterialFieldDetails{}
	}
	return response
}
