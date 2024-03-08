package storage

import (
	"github.com/akhil-is-watching/sainterview-backend/config"
	"github.com/akhil-is-watching/sainterview-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var defaultDB *gorm.DB

func ConnectDB(config *config.Config) {
	db, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})
	if err != nil {
		panic("HumanPal Connection failed")
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	db.Logger = logger.Default.LogMode(logger.Info)

	err = db.AutoMigrate(&models.Avatar{}, &models.Organization{}, &models.Interview{}, &models.InterviewQuestion{})
	if err != nil {
		panic("DB Migrations Failed")
	}

	defaultDB = db
}

func GetDB() *gorm.DB {
	return defaultDB
}
