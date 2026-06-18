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

type UtilisateurControllers struct {
	service *services.UtilisateurService
}

func InitUtilisateurController(service *services.UtilisateurService) *UtilisateurControllers {
	return &UtilisateurControllers{service: service}
}

func readUtilisateurId(r *http.Request) (int, error) {
	return strconv.Atoi(mux.Vars(r)["id"])
}

func (c *UtilisateurControllers) Create(w http.ResponseWriter, r *http.Request) {
	var newUser models.Utilisateur
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "JSON invalide")
		return
	}

	userId, err := c.service.Create(newUser)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := c.service.ReadById(userId)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusCreated, user)
}

func (c *UtilisateurControllers) ReadAll(w http.ResponseWriter, r *http.Request) {
	userList, err := c.service.ReadAll()
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, userList)
}

func (c *UtilisateurControllers) ReadById(w http.ResponseWriter, r *http.Request) {
	idUser, errId := readUtilisateurId(r)
	if errId != nil {
		helper.WriteError(w, http.StatusBadRequest, "Identifiant utilisateur invalide")
		return
	}

	user, err := c.service.ReadById(idUser)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if user.ID == 0 {
		helper.WriteError(w, http.StatusNotFound, "Utilisateur introuvable")
		return
	}

	helper.WriteJSON(w, http.StatusOK, user)
}

func (c *UtilisateurControllers) UpdateById(w http.ResponseWriter, r *http.Request) {
	idUser, errId := readUtilisateurId(r)
	if errId != nil {
		helper.WriteError(w, http.StatusBadRequest, "Identifiant utilisateur invalide")
		return
	}

	var user models.Utilisateur
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "JSON invalide")
		return
	}
	user.ID = idUser

	err := c.service.UpdateById(user)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	updatedUser, err := c.service.ReadById(idUser)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, updatedUser)
}

func (c *UtilisateurControllers) DeleteById(w http.ResponseWriter, r *http.Request) {
	idUser, errId := readUtilisateurId(r)
	if errId != nil {
		helper.WriteError(w, http.StatusBadRequest, "Identifiant utilisateur invalide")
		return
	}

	err := c.service.DeleteById(idUser)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Utilisateur supprimé",
	})
}

func (c *UtilisateurControllers) BanUser(w http.ResponseWriter, r *http.Request) {
	idUser, errId := readUtilisateurId(r)
	if errId != nil {
		helper.WriteError(w, http.StatusBadRequest, "Identifiant utilisateur invalide")
		return
	}

	err := c.service.BanUser(idUser)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Utilisateur banni",
	})
}
