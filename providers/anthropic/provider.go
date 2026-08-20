package anthropic

import (
	"net/http"

	"github.com/imcrazytwkr/feedhub/providers"
	h "github.com/imcrazytwkr/feedhub/providers/anthropic/http"
)

type anthropicProvider struct {
	client *h.AnthropicClient
}

func NewAnthropicProvider(httpClient *http.Client) providers.AnthropicProvider {
	return &anthropicProvider{
		client: h.NewAnthropicClient(httpClient),
	}
}
