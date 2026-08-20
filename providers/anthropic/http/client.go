package http

import (
	"net/http"

	"github.com/imcrazytwkr/feedhub/utils/caches"
	"github.com/imcrazytwkr/feedhub/utils/caches/groupcache"
	"golang.org/x/sync/singleflight"
)

type AnthropicClient struct {
	httpClient   *http.Client
	cache        caches.LRU[string, []byte]
	articleGroup singleflight.Group
}

func NewAnthropicClient(httpClient *http.Client) *AnthropicClient {
	cache, err := groupcache.NewLRU[string, []byte](maxCacheEntries)
	if err != nil {
		panic(err)
	}

	return &AnthropicClient{
		httpClient: httpClient,
		cache:      cache,
	}
}
