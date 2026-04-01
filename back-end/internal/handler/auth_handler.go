package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tuongluong2000/webPhoneStore/back-end/internal/service"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var req LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	token, err := service.Login(req.Email, req.Password)

	if err != nil {
		http.Error(w, "invalid login", 401)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
	})
}
