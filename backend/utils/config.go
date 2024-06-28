package utils

import "github.com/joho/godotenv"

func LoadEnv() {
	err := godotenv.Load()
}