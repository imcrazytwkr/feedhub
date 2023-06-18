package mappers

import (
	"bytes"
	"strconv"
	"strings"
	"time"

	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/imcrazytwkr/feedhub/models"
	m "github.com/imcrazytwkr/feedhub/providers/arknights/models"
	"github.com/imcrazytwkr/feedhub/utils/feedutil"
	"github.com/valyala/fastjson"
)

func PluckEntries(contents *fastjson.Value, language m.Language) []*models.Entry {
	items := contents.GetArray("data", "items")
	if len(items) == 0 {
		return nil
	}

	entries := make([]*models.Entry, len(items))
	i := 0
	for _, item := range items {
		id := parseId(item.Get("id"))
		if len(id) == 0 {
			continue
		}

		entries[i] = &models.Entry{
			Title:     parseText(item.Get("title")),
			Published: parsePublished(item.Get("publishedAt")),
			Link:      generateLink(id, language),
			Content:   parseContent(item.Get("content")),
		}

		i++
	}

	if i < len(entries) {
		entries = entries[:i]
	}

	return feedutil.SortEntries(entries)
}

func parseText(value *fastjson.Value) string {
	if value == nil {
		return ""
	}

	text, err := value.StringBytes()
	if err != nil {
		return ""
	}

	// bytes version works faster and makes less allocations
	return string(bytes.TrimSpace(text))
}

func parseId(value *fastjson.Value) string {
	id := parseText(value)
	if len(id) == 0 {
		return id
	}

	_, err := strconv.Atoi(id)
	if err != nil {
		return ""
	}

	return id
}

func parsePublished(value *fastjson.Value) time.Time {
	text := parseText(value)
	if len(text) == 0 {
		return constants.TimeZero
	}

	timestamp, err := time.Parse(time.RFC3339, text)
	if err != nil {
		return constants.TimeZero
	}

	return timestamp
}

func generateLink(id string, language m.Language) string {
	return hostPrefixes[language] + id
}

func parseContent(value *fastjson.Value) string {
	if value == nil {
		return ""
	}

	array, err := value.Array()
	if err != nil || len(array) == 0 {
		return ""
	}

	// Fast-tracking for the most common variant
	if len(array) == 1 {
		return parseText(array[0].Get("value"))
	}

	builder := strings.Builder{}
	for _, node := range array {
		text := parseText(node.Get("value"))
		if len(text) > 0 {
			builder.WriteString(text)
		}
	}

	if builder.Len() == 0 {
		return ""
	}

	return builder.String()
}
