package services

import (
	"fmt"
	"mathematica-forum/models"
	"mathematica-forum/repositories"
	"strings"
)

type TagService struct {
	tagRepo *repositories.TagRepository
}

func InitTagService(tagRepo *repositories.TagRepository) *TagService {
	return &TagService{tagRepo}
}

func (s *TagService) CreateTag(tag models.Tag) (int, error) {
	if strings.TrimSpace(tag.Nom) == "" {
		return -1, fmt.Errorf("Nom du tag requis")
	}

	if len(tag.Nom) > 50 {
		return -1, fmt.Errorf("Nom du tag trop long")
	}

	id, errCreate := s.tagRepo.CreateTag(tag)
	if errCreate != nil {
		return -1, errCreate
	}

	return id, nil
}

func (s *TagService) GetAllTags() ([]models.Tag, error) {
	tags, errRead := s.tagRepo.ReadAll()
	if errRead != nil {
		return nil, errRead
	}

	return tags, nil
}

func (s *TagService) GetTagById(tagID int) (models.Tag, error) {
	if tagID == 0 {
		return models.Tag{}, fmt.Errorf("ID tag invalide")
	}

	tag, errRead := s.tagRepo.ReadById(tagID)
	if errRead != nil {
		return models.Tag{}, errRead
	}

	return tag, nil
}

func (s *TagService) DeleteTag(tagID int) error {
	if tagID == 0 {
		return fmt.Errorf("ID tag invalide")
	}

	errDelete := s.tagRepo.DeleteTagById(tagID)
	if errDelete != nil {
		return errDelete
	}

	return nil
}

func (s *TagService) GetFilsByTag(tagID int) ([]models.Fil, error) {
	if tagID == 0 {
		return nil, fmt.Errorf("ID tag invalide")
	}

	fils, errRead := s.tagRepo.GetFilsByTag(tagID)
	if errRead != nil {
		return nil, errRead
	}

	return fils, nil
}

func (s *TagService) AddTagToFil(filID int, tagID int) error {
	if filID == 0 || tagID == 0 {
		return fmt.Errorf("IDs invalides")
	}

	errAdd := s.tagRepo.AddTagToFil(filID, tagID)
	if errAdd != nil {
		return errAdd
	}

	return nil
}

func (s *TagService) RemoveTagFromFil(filID int, tagID int) error {
	if filID == 0 || tagID == 0 {
		return fmt.Errorf("IDs invalides")
	}

	errRemove := s.tagRepo.RemoveTagFromFil(filID, tagID)
	if errRemove != nil {
		return errRemove
	}

	return nil
}
