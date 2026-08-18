package models

type ApiError struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
