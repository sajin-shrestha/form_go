package database

import (
	"fmt"

	"github.com/sajin-shrestha/form_go/models"
	"github.com/sajin-shrestha/form_go/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s pprt=%s sslmode=disable",
		utils.GetEnv("DB_HOST"),
		utils.GetEnv("DB_USER"),
		utils.GetEnv("DB_PASSWORD"),
		utils.GetEnv("DB_NAME"),
		utils.GetEnv("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	DB.AutoMigrate(&models.User{})
}
