package services

import (
	"fmt"
	"mathematica-forum/models"
	"mathematica-forum/repositories"
)

type MessageService struct {
	messageRepository *repositories.MessageRepository
}

func InitMessageService(messageRepository *repositories.MessageRepository) *MessageService {
	return &MessageService{messageRepository: messageRepository}
}

func (s *MessageService) Create(message models.Message) (int, error) {
	if message.Contenu == "" || message.UtilisateurID <= 0 || message.FilID <= 0 {
		return -1, fmt.Errorf("Erreur création message - Données manquantes ou invalides")
	}

	if len(message.Contenu) > 5000 {
		return -1, fmt.Errorf("Erreur création message - Contenu trop long (max 5000 caractères)")
	}

	messageId, err := s.messageRepository.CreateMessage(message)
	if err != nil {
		return -1, err
	}

	return messageId, nil
}

func (s *MessageService) ReadAll() ([]models.Message, error) {
	messageList, err := s.messageRepository.ReadAll()
	if err != nil {
		return nil, err
	}

	return messageList, nil
}

func (s *MessageService) ReadById(idMessage int) (models.Message, error) {
	if idMessage <= 0 {
		return models.Message{}, fmt.Errorf("Erreur récupération message - Identifiant invalide : %d", idMessage)
	}

	message, err := s.messageRepository.ReadById(idMessage)
	if err != nil {
		return models.Message{}, err
	}

	return message, nil
}

func (s *MessageService) ReadByFilId(filId int) ([]models.Message, error) {
	if filId <= 0 {
		return nil, fmt.Errorf("Erreur récupération messages - Identifiant fil invalide : %d", filId)
	}

	messages, err := s.messageRepository.ReadByFilId(filId)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (s *MessageService) UpdateById(message models.Message) error {
	if message.ID <= 0 || message.Contenu == "" {
		return fmt.Errorf("Erreur modification message - Données manquantes ou invalides")
	}

	if len(message.Contenu) > 5000 {
		return fmt.Errorf("Erreur modification message - Contenu trop long (max 5000 caractères)")
	}

	return s.messageRepository.UpdateMessageById(message)
}

func (s *MessageService) DeleteById(idMessage int) error {
	if idMessage <= 0 {
		return fmt.Errorf("Erreur suppression message - Identifiant invalide : %d", idMessage)
	}

	return s.messageRepository.DeleteMessageById(idMessage)
}

func (s *MessageService) GetByFilWithPagination(filID int, page int, limit int) ([]models.Message, error) {
	if filID <= 0 {
		return nil, fmt.Errorf("ID fil invalide")
	}

	if page <= 0 {
		page = 1
	}

	if limit <= 0 || limit > 50 {
		limit = 10
	}

	messages, errRead := s.messageRepository.ReadByFilIdPaginated(filID, page, limit)
	if errRead != nil {
		return nil, errRead
	}

	return messages, nil
}
