package configs

import (
	"daijai-service/models/dao"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var database *gorm.DB

func InitDb() *gorm.DB {
	InitConfigFile()
	dsn := viper.GetString("database.path")
	log.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&dao.Project{})
	if err != nil {
		panic("failed to migrate database")
	}

	log.Println("Database migrated")
	return db
}

func GetDBInstance() *gorm.DB {
	if database == nil {
		log.Println("Database is not initialized")
		return InitDb()
	}
	log.Println("GetDBInstance")
	return database
}

func InitConfigFile() {
	// Set the config file name and path
	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	fmt.Println("Successfully loaded config.yaml")

}
