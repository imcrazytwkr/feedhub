package mappers

import (
	"github.com/imcrazytwkr/feedhub/models"
	m "github.com/imcrazytwkr/feedhub/providers/j-novel/models"
)

func ParseEventEntry(event *m.Event) *models.Entry {
	return &models.Entry{
		Title:       event.GetTitle(),
		Published:   event.GetUpdatedAt().AsTime(),
		Link:        "https://j-novel.club/series/" + event.GetSeries().GetSlug(),
		Description: event.GetEventDescription(),
	}
}
