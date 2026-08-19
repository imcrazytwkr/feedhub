package mappers

import (
	"github.com/imcrazytwkr/feedhub/models"
)

func GenerateSiteMeta(language models.Language, category string) *models.Feed {
	feed := &models.Feed{
		Title:       blogsTitle,
		Description: blogsDescription,
		Language:    language.String(),
		Link:        blogsHost,
	}

	if language != models.LanguageEn {
		feed.Link += "/" + language.String()
	}

	feed.Link += blogsPath

	if len(category) > 0 {
		feed.Title += " (" + category + ")"
		feed.Link += "/" + category
	}

	feed.Link += "/"

	return feed
}
