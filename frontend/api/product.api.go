// Package api fournit un client HTTP pour appeler l'API produit externe.
package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"exempleInjection/dto"
)

// ProductApi contient l'URL de base utilisée pour les appels API.
type ProductApi struct {
	baseURL string
}

// InitProductApi initialise le client API avec l'URL de base fournie.
func InitProductApi(baseURL string) *ProductApi {
	return &ProductApi{baseURL: baseURL}
}

// executeRequest envoie la requête HTTP et décode la réponse JSON dans result.
func (api *ProductApi) executeRequest(req *http.Request, result interface{}) (int, error) {
	_client := http.Client{
		Timeout: time.Second * 5,
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := _client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return resp.StatusCode, readErr
	}

	if resp.StatusCode >= 400 {
		return resp.StatusCode, fmt.Errorf("Erreur response - %s", string(bytes.TrimSpace(body)))
	}

	if result != nil {
		if err = json.Unmarshal(body, result); err != nil {
			return resp.StatusCode, fmt.Errorf("Erreur decode données - %s", err.Error())
		}
	}

	return resp.StatusCode, nil
}

// Create envoie une requête POST au point de terminaison /products pour créer un nouveau produit.
func (api *ProductApi) Create(product dto.ProductDto) (int, dto.ProductDto, error) {
	payload, err := json.Marshal(product)
	if err != nil {
		return 0, dto.ProductDto{}, err
	}

	req, err := http.NewRequest(http.MethodPost, api.baseURL+"/products", bytes.NewReader(payload))
	if err != nil {
		return 0, dto.ProductDto{}, err
	}

	var created dto.ProductDto
	_, err = api.executeRequest(req, &created)
	if err != nil {
		return 0, dto.ProductDto{}, err
	}

	return created.Id, created, nil
}

// ReadAll récupère tous les produits depuis l'API distante.
func (api *ProductApi) ReadAll() ([]dto.ProductDto, error) {
	req, err := http.NewRequest(http.MethodGet, api.baseURL+"/products", nil)
	if err != nil {
		return nil, err
	}

	var list []dto.ProductDto
	_, err = api.executeRequest(req, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// ReadById récupère un produit en fonction de son identifiant.
func (api *ProductApi) ReadById(id int) (dto.ProductDto, error) {
	req, err := http.NewRequest(http.MethodGet, api.baseURL+"/products/"+strconv.Itoa(id), nil)
	if err != nil {
		return dto.ProductDto{}, err
	}

	var product dto.ProductDto
	status, err := api.executeRequest(req, &product)
	if err != nil {
		if status == http.StatusNotFound {
			return dto.ProductDto{}, nil
		}
		return dto.ProductDto{}, err
	}

	return product, nil
}

// UpdateById met à jour un produit existant via l'API distante.
func (api *ProductApi) UpdateById(product dto.ProductDto) error {
	payload, err := json.Marshal(product)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, api.baseURL+"/products/"+strconv.Itoa(product.Id), bytes.NewReader(payload))
	if err != nil {
		return err
	}

	status, err := api.executeRequest(req, nil)
	if err != nil {
		if status == http.StatusNotFound {
			return fmt.Errorf("Produit introuvable")
		}
		return err
	}

	return nil
}

// DeleteById supprime un produit via l'API distante.
func (api *ProductApi) DeleteById(id int) error {
	req, err := http.NewRequest(http.MethodDelete, api.baseURL+"/products/"+strconv.Itoa(id), nil)
	if err != nil {
		return err
	}

	status, err := api.executeRequest(req, nil)
	if err != nil {
		if status == http.StatusNotFound {
			return fmt.Errorf("Produit introuvable")
		}
		return err
	}

	return nil
}
