package mappers

import (
	"sort"
	"strconv"

	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/valyala/fastjson"
)

func PluckIllustrationIds(contents *fastjson.Value) ([]int, error) {
	if contents == nil {
		return nil, nil
	}

	err, hasError := processErrorFields(contents)
	if hasError {
		return nil, err
	}

	payload := contents.GetObject(bodyKey, illustsKey)
	if payload == nil || payload.Len() == 0 {
		return nil, nil
	}

	targetLength := payload.Len()
	illustKeys := make([]int, targetLength)
	i := 0

	payload.Visit(func(key []byte, _ *fastjson.Value) {
		illustKey, err := strconv.Atoi(string(key))
		if err != nil {
			return
		}

		illustKeys[i] = illustKey
		i++
	})

	if i < targetLength {
		return nil, constants.ErrorMalformedBody
	}

	sort.Sort(sort.Reverse(sort.IntSlice(illustKeys)))
	return illustKeys, nil
}
