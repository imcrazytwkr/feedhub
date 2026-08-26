package http

import (
	"net/http"

	"github.com/imcrazytwkr/feedhub/utils/caches"
	"github.com/imcrazytwkr/feedhub/utils/caches/groupcache"
	"golang.org/x/sync/singleflight"
)

type CursorClient struct {
	httpClient   *http.Client
	cache        caches.LRU[string, []byte]
	articleGroup singleflight.Group
}

func NewCursorClient(httpClient *http.Client) *CursorClient {
	cache, err := groupcache.NewLRU[string, []byte](maxCacheEntries)
	if err != nil {
		panic(err)
	}

	return &CursorClient{
		httpClient: httpClient,
		cache:      cache,
	}
}
