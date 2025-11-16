package http

import (
	"context"
	"fmt"
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

	uri, err := getNewsUrl(language)
	if err != nil {
		log.Error().Err(err).Msg("failed to generate request URI")
		return nil, models.NewHttpError(http.StatusInternalServerError, nil)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to create request to fetch news feed")
		return nil, models.NewHttpError(http.StatusInternalServerError, nil)
	}

	return httputil.FetchRequest(c.httpClient, req)
}

func getNewsUrl(language m.Language) (string, error) {
	hostPrefix, exists := hostPrefixes[language]
	if !exists {
		return "", fmt.Errorf("mapping for language %s does not exist", language)
	}

	uri, err := url.Parse(hostPrefix)
	if err != nil {
		return "", err
	}

	uri.Path = newsPath
	uri.RawQuery = getNewsQuery(language)

	return uri.String(), nil
}

func getNewsQuery(language m.Language) string {
	query := url.Values{}
	query.Add("lang", language.String())
	query.Add("limit", strconv.Itoa(entryLimit))
	return query.Encode()
}
