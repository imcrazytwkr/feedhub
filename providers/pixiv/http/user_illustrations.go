package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/utils/httputil"
	"github.com/rs/zerolog"
)

func (p *PixivClient) FetchUserIllustrations(ctx context.Context, userId int) ([]byte, error) {
	if userId == 0 {
		return nil, nil
	}

	log := zerolog.Ctx(ctx)

	url := getUserIllustrationsUrl(userId)
	log.Trace().Str("illustration_ids_url", url).Msg("generated illustration ids url")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to create request to fetch illustration ids")
		return nil, models.NewHttpError(http.StatusInternalServerError, nil)
	}

	return httputil.FetchRequest(p.httpClient, req)
}

func getUserIllustrationsUrl(userId int) string {
	return hostPrefix + userPrefix + strconv.Itoa(userId) + latestIllustrationsPrefix
}
