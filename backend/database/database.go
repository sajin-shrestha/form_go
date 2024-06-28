package database

import (
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s pprt=%s sslmode=disable",)
}