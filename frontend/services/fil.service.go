package services

import (
	"encoding/json"
	"forum-frontend/dto"
	"net/http"
)

func GetAllFils() ([]dto.FilResponse, error) {
	resp, err := http.Get(apiBaseURL + "/fils")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var fils []dto.FilResponse
	err = json.NewDecoder(resp.Body).Decode(&fils)
	if err != nil {
		return nil, err
	}

	return fils, nil
}
