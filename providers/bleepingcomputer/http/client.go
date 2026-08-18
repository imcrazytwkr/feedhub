package http

import (
	"net/http"

	"github.com/imcrazytwkr/feedhub/utils/caches"
	"github.com/imcrazytwkr/feedhub/utils/caches/groupcache"
)

// Max lenght of BC news feed * 2
const maxCacheEntries = 30

type BleepingComputerClient struct {
	httpClient *http.Client
	cache      caches.LRU[string, []byte]
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
