package handlers

import (
	"encoding/json"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(*Claims)
	json.NewEncoder(w).Encode(claims)
}
