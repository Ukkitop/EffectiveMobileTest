package database

import (
	"EffectiveMobileTest/internal/database/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func GetDatabaseConnection() (*gorm.DB, error) {
	dsn, exists := os.LookupEnv("DB_CONNECTION")

	if !exists {
		log.Fatal("DB_CONNECTION environment variable not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return db, nil
}

func SetupDatabase() error {
	db, conErr := GetDatabaseConnection()

	if conErr != nil {
		log.Error(conErr)
		return conErr
	}

	err := db.AutoMigrate(&models.DataModel{})
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
