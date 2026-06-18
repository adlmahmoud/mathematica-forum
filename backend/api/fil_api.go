// Package api fournit un client HTTP pour appeler l'API fil externe.
package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"mathematica-forum/models"
)

// FilApi contient l'URL de base utilisée pour les appels API.
type FilApi struct {
	baseURL string
}

// InitFilApi initialise le client API avec l'URL de base fournie.
func InitFilApi(baseURL string) *FilApi {
	return &FilApi{baseURL: baseURL}
}

// executeRequest envoie la requête HTTP et décode la réponse JSON dans result.
func (api *FilApi) executeRequest(req *http.Request, result interface{}) (int, error) {
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

// Create envoie une requête POST au point de terminaison /fils pour créer un nouveau fil.
func (api *FilApi) Create(fil models.Fil) (int, models.Fil, error) {
	payload, err := json.Marshal(fil)
	if err != nil {
		return 0, models.Fil{}, err
	}

	req, err := http.NewRequest(http.MethodPost, api.baseURL+"/fils", bytes.NewReader(payload))
	if err != nil {
		return 0, models.Fil{}, err
	}

	var created models.Fil
	_, err = api.executeRequest(req, &created)
	if err != nil {
		return 0, models.Fil{}, err
	}

	return created.ID, created, nil
}

// ReadAll récupère tous les fils depuis l'API distante.
func (api *FilApi) ReadAll() ([]models.Fil, error) {
	req, err := http.NewRequest(http.MethodGet, api.baseURL+"/fils", nil)
	if err != nil {
		return nil, err
	}

	var list []models.Fil
	_, err = api.executeRequest(req, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// ReadById récupère un fil en fonction de son identifiant.
func (api *FilApi) ReadById(id int) (models.Fil, error) {
	req, err := http.NewRequest(http.MethodGet, api.baseURL+"/fils/"+strconv.Itoa(id), nil)
	if err != nil {
		return models.Fil{}, err
	}

	var fil models.Fil
	status, err := api.executeRequest(req, &fil)
	if err != nil {
		if status == http.StatusNotFound {
			return models.Fil{}, nil
		}
		return models.Fil{}, err
	}

	return fil, nil
}

// UpdateById met à jour un fil existant via l'API distante.
func (api *FilApi) UpdateById(fil models.Fil) error {
	payload, err := json.Marshal(fil)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, api.baseURL+"/fils/"+strconv.Itoa(fil.ID), bytes.NewReader(payload))
	if err != nil {
		return err
	}

	status, err := api.executeRequest(req, nil)
	if err != nil {
		if status == http.StatusNotFound {
			return fmt.Errorf("Fil introuvable")
		}
		return err
	}

	return nil
}

// DeleteById supprime un fil via l'API distante.
func (api *FilApi) DeleteById(id int) error {
	req, err := http.NewRequest(http.MethodDelete, api.baseURL+"/fils/"+strconv.Itoa(id), nil)
	if err != nil {
		return err
	}

	status, err := api.executeRequest(req, nil)
	if err != nil {
		if status == http.StatusNotFound {
			return fmt.Errorf("Fil introuvable")
		}
		return err
	}

	return nil
}
