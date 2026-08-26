package http

import (
	"context"
	"net/http"
	"strings"

	"github.com/imcrazytwkr/feedhub/models"
	m "github.com/imcrazytwkr/feedhub/providers/cursor/models"
	"github.com/imcrazytwkr/feedhub/utils/httputil"
	"github.com/rs/zerolog"
)

func (c *CursorClient) FetchBlog(ctx context.Context, locale m.Locale, topic string) ([]byte, error) {
	return c.fetchPage(ctx, blogURL(locale, topic))
}

func blogURL(locale m.Locale, topic string) string {
	var builder strings.Builder
	builder.WriteString(host)

	if len(locale) > 0 {
		builder.WriteByte('/')
		builder.WriteString(locale.String())
	}

	builder.WriteString(blogPath)

	if len(topic) > 0 {
		builder.WriteString("/topic/")
		builder.WriteString(topic)
	}

	return builder.String()
}

func (c *CursorClient) fetchPage(ctx context.Context, pageURL string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, pageURL, nil)
	if err != nil {
		zerolog.Ctx(ctx).Error().Str("url", pageURL).Err(err).Msg("failed to create request to fetch listing page")
		return nil, models.NewHttpError(http.StatusInternalServerError, nil)
	}

	req.Header = headers.Clone()

	return httputil.FetchRequest(c.httpClient, req)
}
