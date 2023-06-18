package http

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/utils/httputil"
	"github.com/rs/zerolog"
)

func (b *BleepingComputerClient) FetchArticle(ctx context.Context, url string) ([]byte, error) {
	log := zerolog.Ctx(ctx)

	b.cacheMutex.Lock()
	defer b.cacheMutex.Unlock()

	cached, ok := b.cache.Get(url)
	if ok {
		log.Trace().Str("url", url).Msg("cache HIT")
		return cached.([]byte), nil
	}

	log.Trace().Str("url", url).Msg("cache MISS, attempting to query")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	req.Header = http.Header{
		constants.UserAgent:            {headerUserAgent},
		constants.AcceptHeader:         {headerAccept},
		constants.AcceptLanguageHeader: {headerAcceptLanguage},
	}

	if err != nil {
		log.Error().Err(err).Msg("failed to create request to fetch news feed")
		return nil, models.NewHttpError(http.StatusInternalServerError, nil)
	}

	// Sleep between 100 and 500ms between requests, randomness isn't crucial enought
	// to use crypto/rand here
	time.Sleep(time.Duration(rand.Intn(401)+100) * time.Millisecond)

	body, err := httputil.FetchRequest(b.httpClient, req)
	if err != nil {
		return nil, err
	}

	b.cache.Add(url, body)
	return body, nil
}
