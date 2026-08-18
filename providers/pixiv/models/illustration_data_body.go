package models

type IllustrationDataBody struct {
	Works     map[string]IllustrationWork `json:"works"`
	ExtraData ExtraData                   `json:"extraData"`
}
