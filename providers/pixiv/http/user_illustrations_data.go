package http

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/utils/httputil"
	"github.com/rs/zerolog"
)

func (p *PixivClient) FetchUserIllustrationsData(
	ctx context.Context,
	userId int,
	illustIds []int,
) ([]byte, error) {
	if userId == 0 || len(illustIds) == 0 {
		return nil, nil
	}

	log := zerolog.Ctx(ctx)

	url := getUserIllustrationsDataUrl(userId, illustIds)
	log.Trace().Str("illustration_data_url", url).Msg("generated illustration data url")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to create request to fetch illustration data")
		return nil, models.NewHttpError(http.StatusInternalServerError, nil)
	}

	return httputil.FetchRequest(p.httpClient, req)
}

func getUserIllustrationsDataUrl(userIds int, illustIds []int) string {
	builder := strings.Builder{}
	builder.WriteString(hostPrefix)
	builder.WriteString(userPrefix)
	builder.WriteString(strconv.Itoa(userIds))
	builder.WriteString(illustrationsDataPrefix)
	for _, key := range illustIds {
		builder.WriteRune('&')
		builder.WriteString(idsKey)
		builder.WriteString(strconv.Itoa(key))
	}

	return builder.String()
}
