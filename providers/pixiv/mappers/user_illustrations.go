package mappers

import (
	"slices"
	"strconv"

	"github.com/imcrazytwkr/feedhub/constants"
	m "github.com/imcrazytwkr/feedhub/providers/pixiv/models"
)

func PluckIllustrationIds(contents *m.Response[m.IllustrationIDsBody]) ([]int, error) {
	if contents == nil {
		return nil, nil
	}

	err, hasError := processErrorFields(&contents.ApiError)
	if hasError {
		return nil, err
	}

	illusts := contents.Body.Illusts
	targetLength := len(illusts)

	if targetLength == 0 {
		return nil, nil
	}

	illustKeys := make([]int, targetLength)
	i := 0

	for key := range illusts {
		illustKey, err := strconv.Atoi(key)
		if err != nil {
			continue
		}

		illustKeys[i] = illustKey
		i++
	}

	if i < targetLength {
		return nil, constants.ErrorMalformedBody
	}

	slices.Sort(illustKeys)
	slices.Reverse(illustKeys)
	return illustKeys, nil
}
