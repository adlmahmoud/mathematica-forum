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

type TagController struct {
	tagService *services.TagService
}

func InitTagController(tagService *services.TagService) *TagController {
	return &TagController{tagService: tagService}
}

func readTagId(r *http.Request) (int, error) {
	return strconv.Atoi(mux.Vars(r)["tagID"])
}

func (c *TagController) CreateTag(w http.ResponseWriter, r *http.Request) {
	var tag models.Tag
	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "JSON invalide")
		return
	}

	id, err := c.tagService.CreateTag(tag)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	tag.ID = id
	helper.WriteJSON(w, http.StatusCreated, tag)
}

func (c *TagController) GetAllTags(w http.ResponseWriter, r *http.Request) {
	tags, err := c.tagService.GetAllTags()
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, tags)
}

func (c *TagController) GetTagById(w http.ResponseWriter, r *http.Request) {
	id, errId := readTagId(r)
	if errId != nil {
		helper.WriteError(w, http.StatusBadRequest, "ID tag invalide")
		return
	}

	tag, err := c.tagService.GetTagById(id)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if tag.ID == 0 {
		helper.WriteError(w, http.StatusNotFound, "Tag introuvable")
		return
	}

	helper.WriteJSON(w, http.StatusOK, tag)
}

func (c *TagController) DeleteTag(w http.ResponseWriter, r *http.Request) {
	id, errId := readTagId(r)
	if errId != nil {
		helper.WriteError(w, http.StatusBadRequest, "ID tag invalide")
		return
	}

	err := c.tagService.DeleteTag(id)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Tag supprimé",
	})
}

func (c *TagController) GetFilsByTag(w http.ResponseWriter, r *http.Request) {
	id, errId := readTagId(r)
	if errId != nil {
		helper.WriteError(w, http.StatusBadRequest, "ID tag invalide")
		return
	}

	fils, err := c.tagService.GetFilsByTag(id)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, fils)
}

func (c *TagController) AddTagToFil(w http.ResponseWriter, r *http.Request) {
	filID, errFil := strconv.Atoi(mux.Vars(r)["filID"])
	if errFil != nil {
		helper.WriteError(w, http.StatusBadRequest, "ID fil invalide")
		return
	}

	var req struct {
		TagID int `json:"tag_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "JSON invalide")
		return
	}

	errAdd := c.tagService.AddTagToFil(filID, req.TagID)
	if errAdd != nil {
		helper.WriteError(w, http.StatusBadRequest, errAdd.Error())
		return
	}

	helper.WriteJSON(w, http.StatusCreated, map[string]string{
		"message": "Tag ajouté au fil",
	})
}
