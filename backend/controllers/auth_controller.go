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
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}

func (c *AuthControllers) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "JSON invalide")
		return
	}

	token, user, err := c.authService.Login(req.Email, req.Password)
	if err != nil {
		helper.WriteError(w, http.StatusUnauthorized, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, LoginResponse{
		Token: token,
		User: map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"is_admin": user.IsAdmin,
		},
	})
}
