package mappers

import (
	"strings"

	"github.com/imcrazytwkr/feedhub/models"
	m "github.com/imcrazytwkr/feedhub/providers/cursor/models"
)

func BlogSiteMeta(topic string, lang models.Language, description string) *models.Feed {
	if len(description) == 0 {
		description = blogDescription
	}

	return &models.Feed{
		Title:       blogTitleFor(topic),
		Description: description,
		Language:    lang.String(),
		Link:        blogURL(lang, topic),
	}
}

func blogTitleFor(topic string) string {
	if len(topic) == 0 {
		return blogTitle
	}

	return blogTitle + " (" + topic + ")"
}

func blogURL(lang models.Language, topic string) string {
	locale, _ := m.GetLocaleFor(lang)

	var builder strings.Builder
	builder.WriteString(host)

	if len(locale) > 0 {
		builder.WriteByte('/')
		builder.WriteString(locale.String())
	}

	builder.WriteString(blogPath)

	if len(topic) > 0 {
		builder.WriteString("/topic/")
		builder.WriteString(topic)
	}

	return builder.String()
}
