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

type FilControllers struct {
	service *services.FilService
}

func InitFilController(service *services.FilService) *FilControllers {
	return &FilControllers{service: service}
}

func readFilId(r *http.Request) (int, error) {
	return strconv.Atoi(mux.Vars(r)["id"])
}

func (c *FilControllers) Create(w http.ResponseWriter, r *http.Request) {
	var newFil models.Fil
	if err := json.NewDecoder(r.Body).Decode(&newFil); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "JSON invalide")
		return
	}

	filId, err := c.service.Create(newFil)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	fil, err := c.service.ReadById(filId)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusCreated, fil)
}

func (c *FilControllers) ReadAll(w http.ResponseWriter, r *http.Request) {
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

	filList, err := c.service.GetAllWithPagination(page, limit)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if filList == nil {
		filList = []models.Fil{}
	}

	helper.WriteJSON(w, http.StatusOK, filList)
}

func (c *FilControllers) ReadById(w http.ResponseWriter, r *http.Request) {
	idFil, errId := readFilId(r)
	if errId != nil {
		helper.WriteError(w, http.StatusBadRequest, "Identifiant fil invalide")
		return
	}

	fil, err := c.service.ReadById(idFil)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if fil.ID == 0 {
		helper.WriteError(w, http.StatusNotFound, "Fil introuvable")
		return
	}

	helper.WriteJSON(w, http.StatusOK, fil)
}

func (c *FilControllers) UpdateById(w http.ResponseWriter, r *http.Request) {
	idFil, errId := readFilId(r)
	if errId != nil {
		helper.WriteError(w, http.StatusBadRequest, "Identifiant fil invalide")
		return
	}

	var fil models.Fil
	if err := json.NewDecoder(r.Body).Decode(&fil); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "JSON invalide")
		return
	}
	fil.ID = idFil

	err := c.service.UpdateById(fil)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	updatedFil, err := c.service.ReadById(idFil)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, updatedFil)
}

func (c *FilControllers) DeleteById(w http.ResponseWriter, r *http.Request) {
	idFil, errId := readFilId(r)
	if errId != nil {
		helper.WriteError(w, http.StatusBadRequest, "Identifiant fil invalide")
		return
	}

	err := c.service.DeleteById(idFil)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Fil supprimé",
	})
}
