package models

import "time"

type PublicationPost struct {
	Title       string    `json:"title"`
	PublishedAt time.Time `json:"publishedOn"`
	Summary     string    `json:"summary"`
	Slug        Slug      `json:"slug"`
	Tags        []Tag     `json:"subjects"`
}
