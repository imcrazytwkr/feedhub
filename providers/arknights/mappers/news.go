package mappers

import (
	"strconv"
	"strings"
	"time"

	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/imcrazytwkr/feedhub/models"
	m "github.com/imcrazytwkr/feedhub/providers/arknights/models"
	"github.com/imcrazytwkr/feedhub/utils/feedutil"
)

func PluckEntries(contents *m.NewsResponse, language m.Language) []*models.Entry {
	if contents == nil {
		return nil
	}

	items := contents.Data.Items
	if len(items) == 0 {
		return nil
	}

	entries := make([]*models.Entry, len(items))
	i := 0
	for _, item := range items {
		id := parseId(item.Id)
		if len(id) == 0 {
			continue
		}

		entries[i] = &models.Entry{
			// `Id` should be unset so that it's auto-generated in RSS-compatible format
			Title:     strings.TrimSpace(item.Title),
			Published: parsePublished(item.PublishedAt),
			Link:      generateLink(id, language),
			Content:   parseContent(item.Content),
		}

		i++
	}

	if i < len(entries) {
		entries = entries[:i]
	}

	return feedutil.SortEntries(entries)
}

func parseId(value string) string {
	id := strings.TrimSpace(value)
	if len(id) == 0 {
		return id
	}

	_, err := strconv.Atoi(id)
	if err != nil {
		return ""
	}

	return id
}

func parsePublished(value string) time.Time {
	text := strings.TrimSpace(value)
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

func parseContent(nodes []m.ContentNode) string {
	if len(nodes) == 0 {
		return ""
	}

	// Fast-tracking for the most common variant
	if len(nodes) == 1 {
		return strings.TrimSpace(nodes[0].Value)
	}

	builder := strings.Builder{}
	for _, node := range nodes {
		text := strings.TrimSpace(node.Value)
		if len(text) > 0 {
			builder.WriteString(text)
			builder.WriteByte('\n')
		}
	}

	if builder.Len() == 0 {
		return ""
	}

	return builder.String()
}
