package models

type Article struct {
	Title       string `json:"title"`
	Summary     string `json:"summary"`
	PublishedAt Date   `json:"publishedOn"`
	Slug        Slug   `json:"slug"`
}
