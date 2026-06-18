package api

import (
	"net/http"
	"os"
)

func GetAllFils() (*http.Response, error) {
	baseURL := os.Getenv("BASE_URL")

	return http.Get(baseURL + "/fils")
}
