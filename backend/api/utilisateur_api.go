// Package api fournit un client HTTP pour appeler l'API utilisateur externe.
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

// UtilisateurApi contient l'URL de base utilisée pour les appels API.
type UtilisateurApi struct {
	baseURL string
}

// InitUtilisateurApi initialise le client API avec l'URL de base fournie.
func InitUtilisateurApi(baseURL string) *UtilisateurApi {
	return &UtilisateurApi{baseURL: baseURL}
}

// executeRequest envoie la requête HTTP et décode la réponse JSON dans result.
func (api *UtilisateurApi) executeRequest(req *http.Request, result interface{}) (int, error) {
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

// Create envoie une requête POST au point de terminaison /utilisateurs pour créer un nouvel utilisateur.
func (api *UtilisateurApi) Create(user models.Utilisateur) (int, models.Utilisateur, error) {
	payload, err := json.Marshal(user)
	if err != nil {
		return 0, models.Utilisateur{}, err
	}

	req, err := http.NewRequest(http.MethodPost, api.baseURL+"/utilisateurs", bytes.NewReader(payload))
	if err != nil {
		return 0, models.Utilisateur{}, err
	}

	var created models.Utilisateur
	_, err = api.executeRequest(req, &created)
	if err != nil {
		return 0, models.Utilisateur{}, err
	}

	return created.ID, created, nil
}

// ReadAll récupère tous les utilisateurs depuis l'API distante.
func (api *UtilisateurApi) ReadAll() ([]models.Utilisateur, error) {
	req, err := http.NewRequest(http.MethodGet, api.baseURL+"/utilisateurs", nil)
	if err != nil {
		return nil, err
	}

	var list []models.Utilisateur
	_, err = api.executeRequest(req, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// ReadById récupère un utilisateur en fonction de son identifiant.
func (api *UtilisateurApi) ReadById(id int) (models.Utilisateur, error) {
	req, err := http.NewRequest(http.MethodGet, api.baseURL+"/utilisateurs/"+strconv.Itoa(id), nil)
	if err != nil {
		return models.Utilisateur{}, err
	}

	var user models.Utilisateur
	status, err := api.executeRequest(req, &user)
	if err != nil {
		if status == http.StatusNotFound {
			return models.Utilisateur{}, nil
		}
		return models.Utilisateur{}, err
	}

	return user, nil
}

// UpdateById met à jour un utilisateur existant via l'API distante.
func (api *UtilisateurApi) UpdateById(user models.Utilisateur) error {
	payload, err := json.Marshal(user)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, api.baseURL+"/utilisateurs/"+strconv.Itoa(user.ID), bytes.NewReader(payload))
	if err != nil {
		return err
	}

	status, err := api.executeRequest(req, nil)
	if err != nil {
		if status == http.StatusNotFound {
			return fmt.Errorf("Utilisateur introuvable")
		}
		return err
	}

	return nil
}

// DeleteById supprime un utilisateur via l'API distante.
func (api *UtilisateurApi) DeleteById(id int) error {
	req, err := http.NewRequest(http.MethodDelete, api.baseURL+"/utilisateurs/"+strconv.Itoa(id), nil)
	if err != nil {
		return err
	}

	status, err := api.executeRequest(req, nil)
	if err != nil {
		if status == http.StatusNotFound {
			return fmt.Errorf("Utilisateur introuvable")
		}
		return err
	}

	return nil
}
