package services

import (
	"encoding/json"
	"errors"
	"net/http"

	"mathematica-forum/api"
	"mathematica-forum/dto"
)

func GetAllFils() ([]dto.FilResponse, error) {
	resp, err := api.GetAllFils()
	if err != nil {
		return nil, errors.New("impossible de contacter le serveur backend")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("erreur lors de la récupération des discussions")
	}

	var fils []dto.FilResponse
	err = json.NewDecoder(resp.Body).Decode(&fils)
	if err != nil {
		return nil, errors.New("erreur lors de la lecture des données du serveur")
	}

	return fils, nil
}
