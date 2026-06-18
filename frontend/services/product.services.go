// Package services contient la logique métier et les validations avant appel API.
package services

import (
	"exempleInjection/api"
	"exempleInjection/dto"
	"fmt"
)

// ProductService encapsule le client API pour les opérations produit.
type ProductService struct {
	productApi *api.ProductApi
}

// InitProductService initialise le service produit avec le client API externe.
func InitProductService(productApi *api.ProductApi) *ProductService {
	return &ProductService{productApi: productApi}
}

// Create valide les données du produit et appelle l'API pour créer un produit.
func (s *ProductService) Create(product dto.ProductDto) (int, error) {
	if product.Name == "" || product.Description == "" || product.Price < 0 || product.CategorieId < 0 {
		return -1, fmt.Errorf("Erreur ajout produit - Données manquantes ou invalides")
	}

	productId, _, err := s.productApi.Create(product)
	if err != nil {
		return -1, err
	}

	return productId, nil
}

// ReadAll retourne la liste de produits depuis l'API distante.
func (s *ProductService) ReadAll() ([]dto.ProductDto, error) {
	return s.productApi.ReadAll()
}

// ReadById retourne un produit spécifique par son identifiant.
func (s *ProductService) ReadById(idProduct int) (dto.ProductDto, error) {
	if idProduct <= 0 {
		return dto.ProductDto{}, fmt.Errorf("Erreur récupération produit - identifiant invalide : %d", idProduct)
	}

	return s.productApi.ReadById(idProduct)
}

// UpdateById valide et met à jour un produit via l'API.
func (s *ProductService) UpdateById(product dto.ProductDto) error {
	if product.Id <= 0 {
		return fmt.Errorf("Erreur modification produit - Identifiant invalide")
	}
	if product.Name == "" || product.Description == "" || product.Price < 0 || product.CategorieId < 0 {
		return fmt.Errorf("Erreur modification produit - Données manquantes ou invalides")
	}

	return s.productApi.UpdateById(product)
}

// DeleteById supprime un produit en appelant l'API distante.
func (s *ProductService) DeleteById(idProduct int) error {
	if idProduct <= 0 {
		return fmt.Errorf("Erreur suppression produit - Identifiant invalide")
	}

	return s.productApi.DeleteById(idProduct)
}
