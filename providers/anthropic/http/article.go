package http

import (
	"context"
	"net/http"

	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/utils/httputil"
	"github.com/rs/zerolog"
)

func (c *AnthropicClient) FetchArticle(ctx context.Context, articleURL string) ([]byte, error) {
	log := zerolog.Ctx(ctx).With().Str("url", articleURL).Logger()

	cached, ok := c.cache.Get(articleURL)
	if ok {
		log.Trace().Msg("cache HIT")
		return cached, nil
	}

	log.Trace().Msg("cache MISS, attempting to query")

	detached := log.WithContext(context.WithoutCancel(ctx))
	result, err, _ := c.articleGroup.Do(articleURL, func() (any, error) {
		cached, ok := c.cache.Get(articleURL)
		if ok {
			return cached, nil
		}

		return c.fetchArticle(detached, articleURL)
	})

	if err != nil {
		return nil, err
	}

	return result.([]byte), nil
}

func (c *AnthropicClient) fetchArticle(ctx context.Context, articleURL string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, articleURL, nil)
	if err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("failed to create request to fetch article")
		return nil, models.NewHttpError(http.StatusInternalServerError, nil)
	}

	req.Header = headers.Clone()

	body, err := httputil.FetchRequest(c.httpClient, req)
	if err == nil && len(body) > 0 {
		c.cache.Add(articleURL, body)
	}

	return body, err
}
