package models

type SearchResponse struct {
	Items []SearchHit `json:"items"`
}

type SearchHit struct {
	Item Item `json:"item"`
}

type Item struct {
	AdditionalFields AdditionalFields `json:"additionalFields"`
}

type AdditionalFields struct {
	Title        string `json:"title"`
	PostExcerpt  string `json:"postExcerpt"`
	Link         string `json:"link"`
	Contributors string `json:"contributors"`
	CreatedDate  string `json:"createdDate"`
	ModifiedDate string `json:"modifiedDate"`
}
