package http

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/imcrazytwkr/feedhub/models"
	m "github.com/imcrazytwkr/feedhub/providers/aws/models"
	"github.com/imcrazytwkr/feedhub/utils/httputil"
	"github.com/rs/zerolog"
)

func (c *AWSDirectoryClient) SearchBlogPosts(ctx context.Context, locale m.Locale, category string) ([]byte, error) {
	log := zerolog.Ctx(ctx)

	uri, err := getBlogsSearchURL(locale, category)
	if err != nil {
		log.Error().Err(err).Msg("failed to generate request URI")
		return nil, models.NewHttpError(http.StatusInternalServerError, nil)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to create request to fetch blog posts")
		return nil, models.NewHttpError(http.StatusInternalServerError, nil)
	}

	return httputil.FetchRequest(c.httpClient, req)
}

func getBlogsSearchURL(locale m.Locale, category string) (string, error) {
	query := url.Values{}
	query.Set("item.directoryId", directoryIDBlogPosts)
	query.Set("sort_by", sortByCreatedDate)
	query.Set("sort_order", sortOrderDesc)
	query.Set("size", strconv.Itoa(entryLimit))
	query.Set("item.locale", locale.String())

	if len(category) > 0 {
		query.Set("tags.id", categoryTagPrefix+strings.ToLower(category))
	}

	return searchURL + "?" + query.Encode(), nil
}
