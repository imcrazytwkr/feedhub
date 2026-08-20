package models

type Page[S any] struct {
	Sections []S `json:"sections"`
}
