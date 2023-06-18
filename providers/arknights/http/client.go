package http

import (
	"net/http"
)

type ArknightsClient struct {
	httpClient *http.Client
}

func NewArknightsClient(httpClient *http.Client) *ArknightsClient {
	return &ArknightsClient{httpClient}
}
