package http

import (
	"context"
	"net/http"

	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/utils/httputil"
	"github.com/rs/zerolog"
)

func (b *BleepingComputerClient) FetchNewsFeed(ctx context.Context) ([]byte, error) {
	log := zerolog.Ctx(ctx)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, newsFeedUrl, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to create request to fetch news feed")
		return nil, models.NewHttpError(http.StatusInternalServerError, nil)
	}

	req.Header = http.Header{
		constants.UserAgent:            {headerUserAgent},
		constants.AcceptHeader:         {headerAccept},
		constants.AcceptLanguageHeader: {headerAcceptLanguage},
	}

	return httputil.FetchRequest(b.httpClient, req)
}
