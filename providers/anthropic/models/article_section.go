package models

type ArticleSection struct {
	Articles []*Article `json:"articles"`
}
