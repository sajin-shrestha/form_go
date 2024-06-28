package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sajin-shrestha/form_go/database"
	"github.com/sajin-shrestha/form_go/handlers"
	"github.com/sajin-shrestha/form_go/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	database.ConnectDB()

	router := mux.NewRouter()
	router.HandleFunc("/register", handlers.Register).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.HandleFunc("/home", middleware.AuthMiddleware(handlers.Home)).Methods("GET")

	log.Println("server running on port: " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
