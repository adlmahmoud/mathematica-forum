// Si j'ai bien compris le but c'est de recuperer l'api de l'exterieur en utilisant le DTO (je me suis servie de stackoverflow + exemple + une doc)
package api

import (
	"bytes"
	"encoding/json"
	"forum-frontend/dto"
	"net/http"
)

const apiBaseURL = "http://localhost:8080"

func Login(req dto.LoginRequest) (*http.Response, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return http.Post(apiBaseURL+"/login", "application/json", bytes.NewBuffer(jsonData))
}
