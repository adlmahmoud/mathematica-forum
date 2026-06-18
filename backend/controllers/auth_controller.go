package controllers

import (
	"encoding/json"
	"mathematica-forum/helper"
	"mathematica-forum/services"
	"net/http"
)

type AuthControllers struct {
	authService *services.AuthService
}

func InitAuthController(authService *services.AuthService) *AuthControllers {
	return &AuthControllers{authService: authService}
}

type LoginRequest struct {
	Identifiant string `json:"identifiant"`
	Password    string `json:"password"`
}

func (c *AuthControllers) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Requête invalide")
		return
	}

	token, err := c.authService.Login(req.Identifiant, req.Password)
	if err != nil {
		helper.WriteError(w, http.StatusUnauthorized, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]string{
		"token": token,
	})
}
