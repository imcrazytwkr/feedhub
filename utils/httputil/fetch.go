package httputil

import (
	"io"
	"net/http"

	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/rs/zerolog"
)

func FetchRequest(client *http.Client, req *http.Request) ([]byte, error) {
	log := zerolog.Ctx(req.Context())

	res, err := client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("failed to process request")
		return nil, models.NewHttpError(http.StatusServiceUnavailable, nil)
	}

	if res.StatusCode != http.StatusOK {
		return nil, models.NewHttpError(res.StatusCode, nil)
	}

	if req.Method == http.MethodHead {
		return nil, nil
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		log.Warn().Err(err).Msg("failed to read response body")
		return nil, constants.ErrorMalformedBody
	}

	bodySize := int64(len(body))
	if res.ContentLength > -1 && bodySize != res.ContentLength {
		log.Warn().Int64("content_length", res.ContentLength).Int64("body_size", bodySize).Msg("mismatched response size")
		return nil, constants.ErrorMalformedBody
	}

	return body, nil
}
