package http

import (
	"net/http"
)

type TomRigbyClient struct {
	httpClient *http.Client
}

func NewTomRigbyClient(httpClient *http.Client) *TomRigbyClient {
	return &TomRigbyClient{
		httpClient: httpClient,
	}
}
