package http

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/imcrazytwkr/feedhub/models"
	m "github.com/imcrazytwkr/feedhub/providers/arknights/models"
	"github.com/imcrazytwkr/feedhub/utils/httputil"
	"github.com/rs/zerolog"
)

func (c *ArknightsClient) GetNews(ctx context.Context, language m.Language) ([]byte, error) {
	log := zerolog.Ctx(ctx)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, getNewsUrl(language), nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to create request to fetch news feed")
		return nil, models.NewHttpError(http.StatusInternalServerError, nil)
	}

	return httputil.FetchRequest(c.httpClient, req)
}

func getNewsUrl(language m.Language) string {
	uri, _ := url.Parse(hostPrefixes[language])

	uri.Path = newsPath
	uri.RawQuery = getNewsQuery(language)

	return uri.String()
}

func getNewsQuery(language m.Language) string {
	query := url.Values{}
	query.Add("lang", language.String())
	query.Add("limit", strconv.Itoa(entryLimit))
	return query.Encode()
}
