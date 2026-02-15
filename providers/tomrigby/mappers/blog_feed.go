package mappers

import (
	"strings"
	"time"

	"github.com/antchfx/xmlquery"
	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/utils/feedutil"
)

func PluckEntries(contents *xmlquery.Node) ([]*models.Entry, error) {
	items := xmlquery.Find(contents, "//feed/entry")
	if len(items) == 0 {
		return nil, nil
	}

	entries := make([]*models.Entry, len(items))
	for i, item := range items {
		content, err := feedutil.SanitizeContent(item.SelectElement("content"))
		if err != nil {
			return nil, err
		}

		entries[i] = &models.Entry{
			Title:       parseText(item.SelectElement("title")),
			Published:   parseTime(item.SelectElement("published")),
			Updated:     parseTime(item.SelectElement("updated")),
			Author:      parseAuthor(item.SelectElement("author")),
			Link:        parseText(item.SelectElement("link")),
			Description: parseText(item.SelectElement("summary")),
			Content:     content,
		}
	}

	return feedutil.SortEntries(entries), nil
}

func PickSiteMeta(contents *xmlquery.Node) *models.Feed {
	channel := xmlquery.FindOne(contents, "//feed")
	if channel == nil {
		return nil
	}

	return &models.Feed{
		Title:       parseText(channel.SelectElement("title")),
		Description: parseText(channel.SelectElement("subtitle")),
		Published:   parseTime(channel.SelectElement("published")),
		Updated:     parseTime(channel.SelectElement("updated")),
	}
}

func parseAuthor(value *xmlquery.Node) string {
	if value == nil {
		return ""
	}

	name := parseText(value.SelectElement("name"))
	if len(name) > 0 {
		return name
	}

	return strings.TrimSpace(value.InnerText())
}

func parseText(value *xmlquery.Node) string {
	if value == nil {
		return ""
	}

	return strings.TrimSpace(value.InnerText())
}

func parseTime(value *xmlquery.Node) time.Time {
	text := parseText(value)
	if len(text) < 1 {
		return constants.TimeZero
	}

	timestamp, err := time.Parse(time.RFC3339, text)
	if err != nil {
		return constants.TimeZero
	}

	return timestamp
}
