package db

import (
	"daijai-service/models/dao"
	"daijai-service/repositories"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
)

type MaterialFields struct {
	db *gorm.DB
}

func NewMaterialFieldsRepository(db *gorm.DB) repositories.MaterialFields {
	return &MaterialFields{db: db}
}

func (repo MaterialFields) Insert(req dao.MaterialField) error {
	err := repo.db.Create(&req).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo MaterialFields) CheckExisting(req dao.MaterialField) bool {
	var result dao.MaterialField
	err := repo.db.Where("name = ?", req.Name).First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
		return false
	}
	return true
}

func (repo MaterialFields) GetMaterialAll() []dao.MaterialField {
	var response []dao.MaterialField
	err := repo.db.Find(&response).Error
	if err != nil {
		return nil
	}
	return response
}

func (repo MaterialFields) GetMaterialById(id int) dao.MaterialField {
	var response dao.MaterialField
	err := repo.db.Where("id = ?", id).First(&response).Error
	if err != nil {
		log.Println("error db.material_fields:{}", err)
		return dao.MaterialField{}
	}
	return response
}
