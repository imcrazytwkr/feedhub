package models

type PublicationSection struct {
	Title string             `json:"title"`
	Posts []*PublicationPost `json:"posts,omitempty"`
}
