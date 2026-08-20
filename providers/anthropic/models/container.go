package models

type Container[S any] struct {
	Page *Page[S] `json:"page"`
}
