package services

import (
	"fmt"
	"mathematica-forum/models"
	"mathematica-forum/repositories"
)

type FilService struct {
	filRepository *repositories.FilRepository
}

func InitFilService(filRepository *repositories.FilRepository) *FilService {
	return &FilService{filRepository: filRepository}
}

func (s *FilService) Create(fil models.Fil) (int, error) {
	if fil.Titre == "" || fil.UtilisateurID <= 0 {
		return -1, fmt.Errorf("Erreur création fil - Données manquantes ou invalides")
	}

	filId, err := s.filRepository.CreateFil(fil)
	if err != nil {
		return -1, err
	}

	return filId, nil
}

func (s *FilService) ReadAll() ([]models.Fil, error) {
	filList, err := s.filRepository.ReadAll()
	if err != nil {
		return nil, err
	}

	return filList, nil
}

func (s *FilService) ReadById(idFil int) (models.Fil, error) {
	if idFil <= 0 {
		return models.Fil{}, fmt.Errorf("Erreur récupération fil - Identifiant invalide : %d", idFil)
	}

	fil, err := s.filRepository.ReadById(idFil)
	if err != nil {
		return models.Fil{}, err
	}

	return fil, nil
}

func (s *FilService) UpdateById(fil models.Fil) error {
	if fil.ID <= 0 || fil.Titre == "" || fil.Statut == "" {
		return fmt.Errorf("Erreur modification fil - Données manquantes ou invalides")
	}

	return s.filRepository.UpdateFilById(fil)
}

func (s *FilService) DeleteById(idFil int) error {
	if idFil <= 0 {
		return fmt.Errorf("Erreur suppression fil - Identifiant invalide : %d", idFil)
	}

	return s.filRepository.DeleteFilById(idFil)
}
