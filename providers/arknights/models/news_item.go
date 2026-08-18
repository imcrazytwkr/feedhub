package models

type NewsItem struct {
	Id          string        `json:"id"`
	Title       string        `json:"title"`
	PublishedAt string        `json:"publishedAt"`
	Content     []ContentNode `json:"content"`
}
