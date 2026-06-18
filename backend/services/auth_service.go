package services

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"mathematica-forum/config"
	"mathematica-forum/middleware"
	"mathematica-forum/models"
	"mathematica-forum/repositories"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	userRepository *repositories.UtilisateurRepository
}

func InitAuthService(userRepository *repositories.UtilisateurRepository) *AuthService {
	return &AuthService{userRepository: userRepository}
}

func (s *AuthService) Login(email, password string) (string, models.Utilisateur, error) {
	if email == "" || password == "" {
		return "", models.Utilisateur{}, fmt.Errorf("Email et mot de passe requis")
	}

	user, err := s.userRepository.ReadByEmail(email)
	if err != nil {
		return "", models.Utilisateur{}, fmt.Errorf("Utilisateur non trouvé")
	}

	hash := sha512.New()
	hash.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hash.Sum(nil))

	if hashedPassword != user.PasswordHash {
		return "", models.Utilisateur{}, fmt.Errorf("Mot de passe incorrect")
	}

	claims := &middleware.Claims{
		ID:       user.ID,
		Username: user.Username,
		IsAdmin:  user.IsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.GetRequiredEnv("JWT_SECRET")))
	if err != nil {
		return "", models.Utilisateur{}, fmt.Errorf("Erreur génération token")
	}

	return tokenString, user, nil
}
