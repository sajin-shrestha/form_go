package main

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sajin-shrestha/form_go/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	database.ConnectDB()

	router := mux.NewRouter()
	router.
}