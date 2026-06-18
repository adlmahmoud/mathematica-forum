package controllers

import (
	"encoding/json"
	"mathematica-forum/helper"
	"mathematica-forum/models"
	"mathematica-forum/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MessageControllers struct {
	service *services.MessageService
}

func InitMessageController(service *services.MessageService) *MessageControllers {
	return &MessageControllers{service: service}
}

func readMessageId(r *http.Request) (int, error) {
	return strconv.Atoi(mux.Vars(r)["id"])
}

func readFilIdFromRequest(r *http.Request) (int, error) {
	return strconv.Atoi(mux.Vars(r)["filId"])
}

func (c *MessageControllers) Create(w http.ResponseWriter, r *http.Request) {
	var newMessage models.Message
	if err := json.NewDecoder(r.Body).Decode(&newMessage); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "JSON invalide")
		return
	}

	messageId, err := c.service.Create(newMessage)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	message, err := c.service.ReadById(messageId)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusCreated, message)
}

func (c *MessageControllers) ReadAll(w http.ResponseWriter, r *http.Request) {
	messageList, err := c.service.ReadAll()
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, messageList)
}

func (c *MessageControllers) ReadById(w http.ResponseWriter, r *http.Request) {
	idMessage, errId := readMessageId(r)
	if errId != nil {
		helper.WriteError(w, http.StatusBadRequest, "Identifiant message invalide")
		return
	}

	message, err := c.service.ReadById(idMessage)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if message.ID == 0 {
		helper.WriteError(w, http.StatusNotFound, "Message introuvable")
		return
	}

	helper.WriteJSON(w, http.StatusOK, message)
}

func (c *MessageControllers) ReadByFilId(w http.ResponseWriter, r *http.Request) {
	idFil, errId := readFilIdFromRequest(r)
	if errId != nil {
		helper.WriteError(w, http.StatusBadRequest, "Identifiant fil invalide")
		return
	}

	page := 1
	limit := 10

	queryPage := r.URL.Query().Get("page")
	if queryPage != "" {
		if p, err := strconv.Atoi(queryPage); err == nil && p > 0 {
			page = p
		}
	}

	queryLimit := r.URL.Query().Get("limit")
	if queryLimit != "" {
		if l, err := strconv.Atoi(queryLimit); err == nil && l > 0 && l <= 50 {
			limit = l
		}
	}

	messages, err := c.service.GetByFilWithPagination(idFil, page, limit)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if messages == nil {
		messages = []models.Message{}
	}

	helper.WriteJSON(w, http.StatusOK, messages)
}

func (c *MessageControllers) UpdateById(w http.ResponseWriter, r *http.Request) {
	idMessage, errId := readMessageId(r)
	if errId != nil {
		helper.WriteError(w, http.StatusBadRequest, "Identifiant message invalide")
		return
	}

	var message models.Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "JSON invalide")
		return
	}
	message.ID = idMessage

	err := c.service.UpdateById(message)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	updatedMessage, err := c.service.ReadById(idMessage)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, updatedMessage)
}

func (c *MessageControllers) DeleteById(w http.ResponseWriter, r *http.Request) {
	idMessage, errId := readMessageId(r)
	if errId != nil {
		helper.WriteError(w, http.StatusBadRequest, "Identifiant message invalide")
		return
	}

	err := c.service.DeleteById(idMessage)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Message supprimé",
	})
}
