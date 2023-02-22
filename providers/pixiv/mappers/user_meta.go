package mappers

import (
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/valyala/fastjson"
)

func ExtractFeedFromUserMeta(contents *fastjson.Value) *models.Feed {
	if contents == nil {
		return nil
	}

	payload := contents.Get(bodyKey, extraDataKey, metaKey)
	if payload == nil || payload.Type() != fastjson.TypeObject {
		return nil
	}

	title := payload.GetStringBytes(titleKey)
	if len(title) == 0 {
		return nil
	}

	link := payload.GetStringBytes(canonicalKey)
	if len(link) == 0 {
		return nil
	}

	return &models.Feed{
		Title:       string(title),
		Description: string(payload.GetStringBytes(descriptionHeaderKey)),
		Link:        string(link),
	}
}
