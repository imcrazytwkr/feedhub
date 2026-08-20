package mappers

import (
	"github.com/imcrazytwkr/feedhub/models"
)

func EngineeringSiteMeta() *models.Feed {
	return &models.Feed{
		Language:    models.LanguageEn.String(),
		Title:       engineeringTitle,
		Description: engineeringDescription,
		Link:        host + engineeringPath,
	}
}

func NewsSiteMeta() *models.Feed {
	return &models.Feed{
		Language:    models.LanguageEn.String(),
		Title:       newsTitle,
		Description: newsDescription,
		Link:        host + newsPath,
	}
}

func ResearchSiteMeta(team string) *models.Feed {
	feed := &models.Feed{
		Language:    models.LanguageEn.String(),
		Title:       researchTitle,
		Description: researchDescription,
		Link:        host + researchPath,
	}

	if len(team) > 0 {
		feed.Title += " (" + team + ")"
		feed.Link += "/team/" + team
	}

	return feed
}
