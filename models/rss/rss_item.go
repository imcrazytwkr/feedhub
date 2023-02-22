package rss

import (
	"time"

	"github.com/imcrazytwkr/feedhub/models"
)

type RssItem struct {
	Guid        *RssItemGuid `xml:"guid,omitempty"`       // ID used, required!
	Title       string       `xml:"title"`                // Required!
	PubDate     string       `xml:"pubDate,omitempty"`    // max(published, updated)
	Creator     string       `xml:"dc:creator,omitempty"` // Author
	Link        string       `xml:"link"`                 // Required!
	Description string       `xml:"description"`
	Content     *RssContent  `xml:"content:encoded"` // Required!
	Categories  []string     `xml:"category,omitempty"`
}

func NewRssItem(e *models.Entry) *RssItem {
	return &RssItem{
		Guid:        NewRssItemGuid(e),
		Title:       e.Title,
		PubDate:     e.UpdatedString(time.RFC1123Z),
		Creator:     e.Author,
		Link:        e.Link,
		Description: e.Description,
		Content:     NewRssContent(e.Content),
	}
}
