package orm

import (
	"daijai-service/configs"
	"daijai-service/models/dao"
	"gorm.io/gorm"
)

func Save(r *dao.Project) *gorm.DB {
	database := configs.GetDBInstance()
	result := database.Create(&r)
	if result.Error != nil {
		panic(result.Error)
	}
	return database
}
