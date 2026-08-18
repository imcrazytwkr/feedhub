package models

type Response[T any] struct {
	ApiError
	Body T `json:"body"`
}
