package mappers

import (
	"sort"
	"strings"
	"time"

	"github.com/antchfx/xmlquery"
	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/imcrazytwkr/feedhub/models"
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

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Published.After(entries[j].Published)
	})

	return entries
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
	if err != nil {
		return constants.TimeZero
	}

	return timestamp
}
