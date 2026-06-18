package services

import (
	"fmt"
	"mathematica-forum/models"
	"mathematica-forum/repositories"
)

type ReactionService struct {
	reactionRepo *repositories.ReactionRepository
}

func InitReactionService(reactionRepo *repositories.ReactionRepository) *ReactionService {
	return &ReactionService{reactionRepo}
}

func (s *ReactionService) CreateReaction(reaction models.Reaction) (int, error) {
	if reaction.UtilisateurID == 0 || reaction.MessageID == 0 {
		return -1, fmt.Errorf("IDs invalides")
	}

	if reaction.Type != "like" && reaction.Type != "dislike" {
		return -1, fmt.Errorf("Type de réaction invalide")
	}

	existingReaction, errCheck := s.reactionRepo.ReadByUserAndMessage(reaction.UtilisateurID, reaction.MessageID)
	if errCheck != nil {
		return -1, errCheck
	}

	if existingReaction.ID != 0 {
		if existingReaction.Type == reaction.Type {
			return -1, fmt.Errorf("Vous avez déjà réagi avec ce type")
		}

		errDelete := s.reactionRepo.DeleteReactionById(existingReaction.ID)
		if errDelete != nil {
			return -1, errDelete
		}
	}

	id, errCreate := s.reactionRepo.CreateReaction(reaction)
	if errCreate != nil {
		return -1, errCreate
	}

	return id, nil
}

func (s *ReactionService) DeleteReaction(reactionID int) error {
	errDelete := s.reactionRepo.DeleteReactionById(reactionID)
	if errDelete != nil {
		return errDelete
	}

	return nil
}

func (s *ReactionService) GetReactionsByMessage(messageID int) ([]models.Reaction, error) {
	reactions, errRead := s.reactionRepo.ReadByMessageID(messageID)
	if errRead != nil {
		return nil, errRead
	}

	return reactions, nil
}

func (s *ReactionService) GetScore(messageID int) (int, error) {
	likes, errLikes := s.reactionRepo.CountLikesByMessageID(messageID)
	if errLikes != nil {
		return 0, errLikes
	}

	dislikes, errDislikes := s.reactionRepo.CountDislikesByMessageID(messageID)
	if errDislikes != nil {
		return 0, errDislikes
	}

	return likes - dislikes, nil
}
