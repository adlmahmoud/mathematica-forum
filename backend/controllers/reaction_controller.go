package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"mathematica-forum/models"
	"mathematica-forum/services"

	"github.com/gorilla/mux"
)

type ReactionController struct {
	reactionService *services.ReactionService
}

func InitReactionController(reactionService *services.ReactionService) *ReactionController {
	return &ReactionController{reactionService}
}

type CreateReactionRequest struct {
	Type string `json:"type"`
}

type ReactionResponse struct {
	ID            int    `json:"id"`
	Type          string `json:"type"`
	UtilisateurID int    `json:"utilisateur_id"`
	MessageID     int    `json:"message_id"`
}

type ScoreResponse struct {
	Score int `json:"score"`
}

func (c *ReactionController) CreateReaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	messageID, errConvert := strconv.Atoi(vars["messageID"])
	if errConvert != nil {
		http.Error(w, "ID message invalide", http.StatusBadRequest)
		return
	}

	userIDStr := r.Header.Get("X-User-ID")
	userID, errParse := strconv.Atoi(userIDStr)
	if errParse != nil {
		http.Error(w, "Authentification requise", http.StatusUnauthorized)
		return
	}

	var req CreateReactionRequest
	errDecode := json.NewDecoder(r.Body).Decode(&req)
	if errDecode != nil {
		http.Error(w, "JSON invalide", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(req.Type) == "" {
		http.Error(w, "Type requis", http.StatusBadRequest)
		return
	}

	reaction := models.Reaction{
		Type:          req.Type,
		UtilisateurID: userID,
		MessageID:     messageID,
	}

	id, errCreate := c.reactionService.CreateReaction(reaction)
	if errCreate != nil {
		http.Error(w, errCreate.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func (c *ReactionController) DeleteReaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reactionID, errConvert := strconv.Atoi(vars["reactionID"])
	if errConvert != nil {
		http.Error(w, "ID réaction invalide", http.StatusBadRequest)
		return
	}

	errDelete := c.reactionService.DeleteReaction(reactionID)
	if errDelete != nil {
		http.Error(w, errDelete.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Réaction supprimée"})
}

func (c *ReactionController) GetReactionsByMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	messageID, errConvert := strconv.Atoi(vars["messageID"])
	if errConvert != nil {
		http.Error(w, "ID message invalide", http.StatusBadRequest)
		return
	}

	reactions, errRead := c.reactionService.GetReactionsByMessage(messageID)
	if errRead != nil {
		http.Error(w, errRead.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reactions)
}

func (c *ReactionController) GetScore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	messageID, errConvert := strconv.Atoi(vars["messageID"])
	if errConvert != nil {
		http.Error(w, "ID message invalide", http.StatusBadRequest)
		return
	}

	score, errScore := c.reactionService.GetScore(messageID)
	if errScore != nil {
		http.Error(w, errScore.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ScoreResponse{Score: score})
}
