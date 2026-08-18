package models

type IllustrationWork struct {
	Title       string `json:"title"`
	UserName    string `json:"userName"`
	UserID      string `json:"userId"`
	CreateDate  string `json:"createDate"`
	UpdateDate  string `json:"updateDate"`
	PageCount   int    `json:"pageCount"`
	URL         string `json:"url"`
	Description string `json:"description"`
}
