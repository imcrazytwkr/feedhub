package mappers

import (
	"strings"
	"time"

	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/imcrazytwkr/feedhub/models"
	m "github.com/imcrazytwkr/feedhub/providers/aws/models"
	"github.com/imcrazytwkr/feedhub/utils/feedutil"
)

func PluckEntries(contents *m.SearchResponse) []*models.Entry {
	if contents == nil {
		return nil
	}

	items := contents.Items
	if len(items) == 0 {
		return nil
	}

	entries := make([]*models.Entry, len(items))
	i := 0
	for _, hit := range items {
		fields := hit.Item.AdditionalFields
		title := strings.TrimSpace(fields.Title)
		link := strings.TrimSpace(fields.Link)
		if len(title) == 0 || len(link) == 0 {
			continue
		}

		entries[i] = &models.Entry{
			Title:       title,
			Published:   parseTimestamp(fields.CreatedDate),
			Updated:     parseTimestamp(fields.ModifiedDate),
			Author:      strings.TrimSpace(fields.Contributors),
			Link:        link,
			Description: strings.TrimSpace(fields.PostExcerpt),
		}
		i++
	}

	if i == 0 {
		return nil
	}

	if i < len(entries) {
		entries = entries[:i]
	}

	return feedutil.SortEntries(entries)
}

func parseTimestamp(value string) time.Time {
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
