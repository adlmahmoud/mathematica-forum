package services

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"mathematica-forum/models"
	"mathematica-forum/repositories"
	"unicode"
)

type UtilisateurService struct {
	userRepository *repositories.UtilisateurRepository
}

func InitUtilisateurService(userRepository *repositories.UtilisateurRepository) *UtilisateurService {
	return &UtilisateurService{userRepository: userRepository}
}

func (s *UtilisateurService) Create(user models.Utilisateur) (int, error) {
	if user.Username == "" || user.Email == "" || len(user.PasswordHash) < 12 {
		return -1, fmt.Errorf("Erreur inscription - Données manquantes ou mot de passe trop court")
	}

	var hasUpper, hasSpecial bool
	for _, r := range user.PasswordHash {
		if unicode.IsUpper(r) {
			hasUpper = true
		}
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			hasSpecial = true
		}
	}

	if !hasUpper || !hasSpecial {
		return -1, fmt.Errorf("Erreur inscription - Le mot de passe doit contenir au moins une majuscule et un caractère spécial")
	}

	hash := sha512.New()
	hash.Write([]byte(user.PasswordHash))
	user.PasswordHash = hex.EncodeToString(hash.Sum(nil))

	userId, err := s.userRepository.CreateUtilisateur(user)
	if err != nil {
		return -1, err
	}

	return userId, nil
}

func (s *UtilisateurService) ReadAll() ([]models.Utilisateur, error) {
	userList, err := s.userRepository.ReadAll()
	if err != nil {
		return nil, err
	}

	return userList, nil
}

func (s *UtilisateurService) ReadById(idUser int) (models.Utilisateur, error) {
	if idUser <= 0 {
		return models.Utilisateur{}, fmt.Errorf("Erreur récupération utilisateur - Identifiant invalide : %d", idUser)
	}

	user, err := s.userRepository.ReadById(idUser)
	if err != nil {
		return models.Utilisateur{}, err
	}

	return user, nil
}

func (s *UtilisateurService) UpdateById(user models.Utilisateur) error {
	if user.ID <= 0 {
		return fmt.Errorf("Erreur modification utilisateur - Identifiant invalide")
	}

	return s.userRepository.UpdateUtilisateurById(user)
}

func (s *UtilisateurService) DeleteById(idUser int) error {
	if idUser <= 0 {
		return fmt.Errorf("Erreur suppression utilisateur - Identifiant invalide : %d", idUser)
	}

	return s.userRepository.DeleteUtilisateurById(idUser)
}

func (s *UtilisateurService) BanUser(userID int) error {
	if userID <= 0 {
		return fmt.Errorf("ID utilisateur invalide")
	}

	user, errRead := s.userRepository.ReadById(userID)
	if errRead != nil {
		return errRead
	}

	if user.ID == 0 {
		return fmt.Errorf("Utilisateur non trouvé")
	}

	user.IsBanni = true
	errUpdate := s.userRepository.UpdateUtilisateurById(user)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}
