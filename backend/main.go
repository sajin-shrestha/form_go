package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sajin-shrestha/form_go/database"

	customHandlers "github.com/sajin-shrestha/form_go/handlers"
	"github.com/sajin-shrestha/form_go/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	database.ConnectDB()

	corsOptions := handlers.AllowedOrigins([]string{"http://localhost:5173"})
	corsHandler := handlers.CORS(corsOptions)

	router := mux.NewRouter()
	router.HandleFunc("/register", customHandlers.Register).Methods("POST")
	router.HandleFunc("/login", customHandlers.Login).Methods("POST")
	router.HandleFunc("/delete/{id}", middleware.AuthMiddleware(customHandlers.DeleteUser)).Methods("DELETE")
	router.HandleFunc("/home", middleware.AuthMiddleware(customHandlers.Home)).Methods("GET")

	log.Println("server running on port: " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), corsHandler(router)))
}
