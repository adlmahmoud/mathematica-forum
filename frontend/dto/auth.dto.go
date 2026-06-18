// Package dto contient les structures de données utilisées pour les échanges entre services et templates.
// C'est comme si que en fait un contrat entre le backend et le frontend
package dto

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Identifiant string `json:"identifiant"`
	Password    string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
