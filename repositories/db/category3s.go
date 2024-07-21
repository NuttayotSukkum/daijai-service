package db

import (
	"daijai-service/models/dao"
	"daijai-service/repositories"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
)

type Category3 struct {
	db *gorm.DB
}

func NewCategory3Repository(db *gorm.DB) repositories.Category3 {
	return &Category3{db: db}
}

func (repo Category3) Insert(rq dao.Category3) error {
	err := repo.db.Create(&rq).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo Category3) GetAllCategory3() ([]dao.Category3, error) {
	var listCategory3 []dao.Category3
	err := repo.db.Find(&listCategory3).Error
	if err != nil {
		return nil, err
	}
	return listCategory3, err
}

func (repo *Category3) CheckCat3isExist(req dao.Category3) bool {
	var response dao.Category3
	log.Println("Checking existence for:", req.Name, req.Code)

	err := repo.db.Where("name = ? AND code = ?", req.Name, req.Code).First(&response).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("Category3 not found")
			return false
		}
		log.Printf("Error checking Category3 existence")
		return false
	}
	log.Println("Category3 exists: user is existing")
	return true
}

func (repo *Category3) GetCat3ById(id int) (dao.Category3, error) {
	var response dao.Category3
	err := repo.db.Where("id = ?", id).First(&response)
	if err != nil {
		return response, err.Error
	}
	return response, nil
}

func (repo *Category3) FindCat3ById(id int) (dao.Category3, error) {
	var res dao.Category3
	err := repo.db.Where("id = ?", id).First(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Category3 not found")
			return res, nil
		}
		log.Printf("Error finding Category3 by ID: %v", err)
		return res, err
	}
	return res, nil
}
