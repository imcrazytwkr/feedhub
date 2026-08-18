package mappers

import (
	"errors"

	m "github.com/imcrazytwkr/feedhub/providers/pixiv/models"
)

func processErrorFields(apiError *m.ApiError) (error, bool) {
	if apiError == nil || !apiError.Error {
		return nil, false
	}

	if len(apiError.Message) == 0 {
		return nil, true
	}

	return errors.New(apiError.Message), true
}
