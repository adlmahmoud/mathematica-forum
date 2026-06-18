// Package controllers contient la logique de présentation pour les pages produits.
package controllers

import (
	models "exempleInjection/dto"
	"exempleInjection/services"
	"exempleInjection/templates"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ProductControllers gère les actions web liées aux produits.
type ProductControllers struct {
	service  *services.ProductService
	template *templates.TemplateManager
}

// InitProductController crée un contrôleur produit avec son service et son moteur de templates.
func InitProductController(service *services.ProductService, template *templates.TemplateManager) *ProductControllers {
	return &ProductControllers{service: service, template: template}
}

// CreateForm affiche le formulaire de création de produit.
func (c *ProductControllers) CreateForm(w http.ResponseWriter, r *http.Request) {
	c.template.RenderTemplate(w, r, "product.create", nil)
}

// Create traite la soumission du formulaire de création et redirige vers la page produit créée.
func (c *ProductControllers) Create(w http.ResponseWriter, r *http.Request) {
	convPrice, convPriceErr := strconv.ParseFloat(r.FormValue("price"), 32)
	convCategorie, convCategorieErr := strconv.Atoi(r.FormValue("categorie"))
	if convPriceErr != nil || convCategorieErr != nil {
		http.Error(w, "Erreur - Des données sont manquantes ou bien invalide", http.StatusBadRequest)
		return
	}

	newProduct := models.ProductDto{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Price:       float32(convPrice),
		CategorieId: convCategorie,
	}

	productId, productErr := c.service.Create(newProduct)
	if productErr != nil {
		http.Error(w, productErr.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/product/%d", productId), http.StatusSeeOther)
}

// DisplayList affiche la liste des produits.
func (c *ProductControllers) DisplayList(w http.ResponseWriter, r *http.Request) {
	productList, productErr := c.service.ReadAll()
	if productErr != nil {
		http.Error(w, productErr.Error(), http.StatusInternalServerError)
		return
	}
	c.template.RenderTemplate(w, r, "product.list", productList)
}

// DisplayById affiche les détails d'un produit identifié par son ID.
func (c *ProductControllers) DisplayById(w http.ResponseWriter, r *http.Request) {
	idProduct, idProductErr := strconv.Atoi(mux.Vars(r)["id"])
	if idProductErr != nil {
		http.Error(w, "Erreur - Identifiant produit invalide", http.StatusBadRequest)
		return
	}

	product, productErr := c.service.ReadById(idProduct)
	if productErr != nil {
		http.Error(w, productErr.Error(), http.StatusInternalServerError)
		return
	}

	c.template.RenderTemplate(w, r, "product", product)
}
