package http

import (
	"context"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/utils/httputil"
	"github.com/rs/zerolog"
)

func (b *BleepingComputerClient) FetchArticle(ctx context.Context, url string) ([]byte, error) {
	log := zerolog.Ctx(ctx).With().Str("url", url).Logger()

	cached, ok := b.cache.Get(url)
	if ok {
		log.Trace().Msg("cache HIT")
		return cached, nil
	}

	log.Trace().Msg("cache MISS, attempting to query")

	detached := log.WithContext(context.WithoutCancel(ctx))
	result, err, _ := b.articleGroup.Do(url, func() (any, error) {
		cached, ok := b.cache.Get(url)
		if ok {
			return cached, nil
		}

		return b.fetchArticle(detached, url)
	})

	if err != nil {
		return nil, err
	}

	return result.([]byte), nil
}

func (b *BleepingComputerClient) fetchArticle(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("failed to create request to fetch news feed")
		return nil, models.NewHttpError(http.StatusInternalServerError, nil)
	}

	req.Header = headers.Clone()

	// Sleep between 100 and 500ms between requests, randomness isn't crucial enought
	// to use crypto/rand here
	time.Sleep(time.Duration(rand.IntN(401)+100) * time.Millisecond)

	body, err := httputil.FetchRequest(b.httpClient, req)
	if err == nil && len(body) > 0 {
		b.cache.Add(url, body)
	}

	return body, err
}
