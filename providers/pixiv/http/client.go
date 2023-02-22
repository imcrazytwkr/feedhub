package http

import "net/http"

type PixivClient struct {
	httpClient *http.Client
}

func NewPixivClient(httpClient *http.Client) *PixivClient {
	return &PixivClient{httpClient}
}
