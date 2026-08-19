package mappers

import (
	"strings"

	"github.com/imcrazytwkr/feedhub/models"
)

func GenerateSiteMeta(language models.Language) *models.Feed {
	return &models.Feed{
		Title:       feedTitles[language],
		Description: feedDescriptions[language],
		Language:    language.String(),
		Link:        strings.TrimRight(hostPrefixes[language], "/"),
	}
}
