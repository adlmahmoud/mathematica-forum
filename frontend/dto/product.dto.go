// Package dto contient les structures de données utilisées pour les échanges entre services et templates.
package dto

// ProductDto représente un produit dans les échanges JSON et les templates.
type ProductDto struct {
	Id          int     `json:"id"`
	Name        string  `json:"nom"`
	Description string  `json:"description"`
	Price       float32 `json:"prix"`
	CategorieId int     `json:"categorie_id"`
	CreateAt    string  `json:"date_ajout"`
}
