package cursor

import (
	"net/http"

	"github.com/imcrazytwkr/feedhub/providers"
	h "github.com/imcrazytwkr/feedhub/providers/cursor/http"
)

type cursorProvider struct {
	client *h.CursorClient
}

func NewCursorProvider(httpClient *http.Client) providers.CursorProvider {
	return &cursorProvider{
		client: h.NewCursorClient(httpClient),
	}
}
