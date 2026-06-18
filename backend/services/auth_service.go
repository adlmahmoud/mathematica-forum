package services

import (
	"crypto/sha512"
	"database/sql"
	"encoding/hex"
	"fmt"
	"mathematica-forum/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	db *sql.DB
}

func InitAuthService(db *sql.DB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) Login(identifiant, password string) (string, error) {
	var user models.Utilisateur
	query := "SELECT id_utilisateur, mot_de_passe_hash, sel, is_banni FROM utilisateur WHERE nom_utilisateur = ? OR email = ?;"

	err := s.db.QueryRow(query, identifiant, identifiant).Scan(&user.ID, &user.PasswordHash, &user.Sel, &user.IsBanni)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("Identifiants incorrects")
		}
		return "", err
	}

	if user.IsBanni {
		return "", fmt.Errorf("Ce compte a été banni")
	}

	hash := sha512.New()
	hash.Write([]byte(password + user.Sel))
	hashedPassword := hex.EncodeToString(hash.Sum(nil))

	if hashedPassword != user.PasswordHash {
		return "", fmt.Errorf("Identifiants incorrects")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "secret_de_secours_a_changer"
	}

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
