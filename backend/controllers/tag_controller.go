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

type TagController struct {
	tagService *services.TagService
}

func InitTagController(tagService *services.TagService) *TagController {
	return &TagController{tagService}
}

type CreateTagRequest struct {
	Nom string `json:"nom"`
}

type TagResponseItem struct {
	ID  int    `json:"id"`
	Nom string `json:"nom"`
}

func (c *TagController) CreateTag(w http.ResponseWriter, r *http.Request) {
	var req CreateTagRequest
	errDecode := json.NewDecoder(r.Body).Decode(&req)
	if errDecode != nil {
		http.Error(w, "JSON invalide", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(req.Nom) == "" {
		http.Error(w, "Nom du tag requis", http.StatusBadRequest)
		return
	}

	tag := models.Tag{Nom: req.Nom}
	id, errCreate := c.tagService.CreateTag(tag)
	if errCreate != nil {
		http.Error(w, errCreate.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func (c *TagController) GetAllTags(w http.ResponseWriter, r *http.Request) {
	tags, errRead := c.tagService.GetAllTags()
	if errRead != nil {
		http.Error(w, errRead.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tags)
}

func (c *TagController) GetTagById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tagID, errConvert := strconv.Atoi(vars["tagID"])
	if errConvert != nil {
		http.Error(w, "ID tag invalide", http.StatusBadRequest)
		return
	}

	tag, errRead := c.tagService.GetTagById(tagID)
	if errRead != nil {
		http.Error(w, errRead.Error(), http.StatusInternalServerError)
		return
	}

	if tag.ID == 0 {
		http.Error(w, "Tag non trouvé", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tag)
}

func (c *TagController) GetFilsByTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tagID, errConvert := strconv.Atoi(vars["tagID"])
	if errConvert != nil {
		http.Error(w, "ID tag invalide", http.StatusBadRequest)
		return
	}

	filIDs, errRead := c.tagService.GetFilsByTag(tagID)
	if errRead != nil {
		http.Error(w, errRead.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]int{"fil_ids": filIDs})
}

func (c *TagController) DeleteTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tagID, errConvert := strconv.Atoi(vars["tagID"])
	if errConvert != nil {
		http.Error(w, "ID tag invalide", http.StatusBadRequest)
		return
	}

	errDelete := c.tagService.DeleteTag(tagID)
	if errDelete != nil {
		http.Error(w, errDelete.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Tag supprimé"})
}

func (c *TagController) AddTagToFil(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filID, errConvert := strconv.Atoi(vars["filID"])
	if errConvert != nil {
		http.Error(w, "ID fil invalide", http.StatusBadRequest)
		return
	}

	var req map[string]int
	errDecode := json.NewDecoder(r.Body).Decode(&req)
	if errDecode != nil {
		http.Error(w, "JSON invalide", http.StatusBadRequest)
		return
	}

	tagID, ok := req["tag_id"]
	if !ok {
		http.Error(w, "tag_id manquant", http.StatusBadRequest)
		return
	}

	errAdd := c.tagService.AddTagToFil(filID, tagID)
	if errAdd != nil {
		http.Error(w, errAdd.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Tag ajouté"})
}
