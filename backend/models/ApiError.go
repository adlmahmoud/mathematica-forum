package models

type ApiError struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}
