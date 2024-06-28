package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/sajin-shrestha/form_go/models"
)


var jwtKey = []byte("my_secret_jwt_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, err := 
}