package http

import (
	"net/http"

	"github.com/imcrazytwkr/feedhub/utils/caches"
	"github.com/imcrazytwkr/feedhub/utils/caches/groupcache"
	"golang.org/x/sync/singleflight"
)

// Max lenght of BC news feed * 2
const maxCacheEntries = 30

type BleepingComputerClient struct {
	httpClient   *http.Client
	cache        caches.LRU[string, []byte]
	articleGroup singleflight.Group
}

func NewBleepingComputerClient(httpClient *http.Client) *BleepingComputerClient {
	cache, err := groupcache.NewLRU[string, []byte](maxCacheEntries)
	if err != nil {
		// Development error, panic early
		panic(err)
	}

	return &BleepingComputerClient{
		httpClient: httpClient,
		cache:      cache,
	}
}
