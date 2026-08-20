package http

import (
	"context"
	"net/http"

	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/utils/httputil"
	"github.com/rs/zerolog"
)

func (c *AnthropicClient) FetchNews(ctx context.Context) ([]byte, error) {
	return c.fetchPage(ctx, newsURL)
}

func (c *AnthropicClient) FetchEngineering(ctx context.Context) ([]byte, error) {
	return c.fetchPage(ctx, engineeringURL)
}

func (c *AnthropicClient) FetchResearch(ctx context.Context) ([]byte, error) {
	return c.fetchPage(ctx, researchURL)
}

func (c *AnthropicClient) fetchPage(ctx context.Context, pageURL string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, pageURL, nil)
	if err != nil {
		zerolog.Ctx(ctx).Error().Str("url", pageURL).Err(err).Msg("failed to create request to fetch listing page")
		return nil, models.NewHttpError(http.StatusInternalServerError, nil)
	}

	req.Header = headers.Clone()

	return httputil.FetchRequest(c.httpClient, req)
}
