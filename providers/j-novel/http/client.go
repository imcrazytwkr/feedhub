package http

import (
	"net/http"
)

type JNovelClient struct {
	httpClient *http.Client
}

func NewJNovelClient(httpClient *http.Client) *JNovelClient {
	return &JNovelClient{httpClient}
}
