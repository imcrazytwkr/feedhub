package mappers

import (
	"strings"
	"time"

	"github.com/antchfx/xmlquery"
	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/utils/feedutil"
)

func PluckItems(contents *xmlquery.Node) []*models.Entry {
	items := xmlquery.Find(contents, "//channel/item")
	if len(items) == 0 {
		return nil
	}

	entries := make([]*models.Entry, len(items))
	for i, item := range items {
		entries[i] = &models.Entry{
			Title:     parseText(item.SelectElement("title")),
			Published: parsePublished(item.SelectElement("pubDate")),
			Author:    parseText(item.SelectElement("dc:creator")),
			Link:      parseText(item.SelectElement("link")),
		}
	}

	return feedutil.SortEntries(entries)
}

func PickSiteMeta(contents *xmlquery.Node) *models.Feed {
	channel := xmlquery.FindOne(contents, "//channel")
	if channel == nil {
		return nil
	}

	return &models.Feed{
		Title:       parseText(channel.SelectElement("title")),
		Description: parseText(channel.SelectElement("description")),
		Language:    parseText(channel.SelectElement("language")),
		Published:   parsePublished(channel.SelectElement("pubDate")),
	}
}

func parseText(value *xmlquery.Node) string {
	if value == nil {
		return ""
	}

	return strings.TrimSpace(value.InnerText())
}

func parsePublished(value *xmlquery.Node) time.Time {
	text := parseText(value)
	if len(text) < 1 {
		return constants.TimeZero
	}

	timestamp, err := time.Parse(time.RFC1123Z, text)
	if err == nil {
		return timestamp
	}

	timestamp, err = time.Parse(time.RFC1123, text)
	if err == nil {
		return timestamp
	}

	return constants.TimeZero
}
