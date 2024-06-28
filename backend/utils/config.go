package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("error loading .env file")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}