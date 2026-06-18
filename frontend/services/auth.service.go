// Le but de se fichier et de recuperer les donners en utilisant le dto, pour permettre au controllers d'afficher et de charger les donnees au templates
package services

import (
	"encoding/json"
	"errors"
	"mathematica-forum/api"
	"mathematica-forum/dto"
	"net/http"
)

func Login(req dto.LoginRequest) (string, error) {
	resp, err := api.Login(req)
	if err != nil {
		return "", errors.New("impossible de contacter le serveur backend")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return "", errors.New("identifiant ou mot de passe incorrect")
	} else if resp.StatusCode != http.StatusOK {
		return "", errors.New("une erreur inattendue est survenue sur le serveur")
	}

	var authResp dto.AuthResponse
	err = json.NewDecoder(resp.Body).Decode(&authResp)
	if err != nil {
		return "", errors.New("erreur lors de la lecture des données du serveur")
	}

	return authResp.Token, nil
}
