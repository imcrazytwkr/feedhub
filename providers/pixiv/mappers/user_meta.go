package mappers

import (
	"github.com/imcrazytwkr/feedhub/models"
	m "github.com/imcrazytwkr/feedhub/providers/pixiv/models"
)

func ExtractFeedFromUserMeta(meta *m.UserMeta) *models.Feed {
	if meta == nil {
		return nil
	}

	if len(meta.Title) == 0 {
		return nil
	}

	if len(meta.Canonical) == 0 {
		return nil
	}

	return &models.Feed{
		Title:       meta.Title,
		Description: meta.DescriptionHeader,
		Link:        meta.Canonical,
	}
}
