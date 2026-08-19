package http

import (
	"net/http"
)

type AWSDirectoryClient struct {
	httpClient *http.Client
}

func NewAWSDirectoryClient(httpClient *http.Client) *AWSDirectoryClient {
	return &AWSDirectoryClient{httpClient}
}
