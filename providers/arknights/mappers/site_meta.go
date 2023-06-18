package mappers

import (
	"strings"

	"github.com/imcrazytwkr/feedhub/models"
	m "github.com/imcrazytwkr/feedhub/providers/arknights/models"
)

func GenerateSiteMeta(language m.Language) *models.Feed {
	return &models.Feed{
		Title:       feedTitles[language],
		Description: feedDescriptions[language],
		Language:    language.String(),
		Link:        strings.TrimRight(hostPrefixes[language], "/"),
	}
}
