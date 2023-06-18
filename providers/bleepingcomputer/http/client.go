package http

import (
	"net/http"
	"sync"

	"github.com/golang/groupcache/lru"
)

// Max lenght of BC news feed * 2
const maxCacheEntries = 30

type BleepingComputerClient struct {
	httpClient *http.Client
	cache      *lru.Cache
	cacheMutex sync.Mutex
}

func NewPixivClient(httpClient *http.Client) *BleepingComputerClient {
	return &BleepingComputerClient{
		httpClient: httpClient,
		cache:      lru.New(maxCacheEntries),
	}
}
