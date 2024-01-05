package utils

import (
	"fmt"
	"go-fiber/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() error {
	var err error
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Config("DB_HOST"), Config("DB_PORT"), Config("DB_USERNAME"), Config("DB_USER_PASSWORD"), Config("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error in connect the DB %v", err)
		return nil
	}

	if DB.Error != nil {
		log.Fatalln("Any Error in connect the DB " + err.Error())
		return nil
	}

	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	DB.AutoMigrate(&models.Book{})

	log.Println("DB connected")

	return nil
}
