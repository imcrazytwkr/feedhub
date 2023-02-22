package rss

import (
	"encoding/xml"
	"time"

	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/models/atom"
)

const RssMime = "application/rss+xml"
const rssGenerator = "FeedHub/Rss2"
const defaultLanguage = "en-us"

type RssFeed struct {
	XMLName       xml.Name       `xml:"channel"`
	Generator     string         `xml:"generator"`
	Language      string         `xml:"language"`
	Title         string         `xml:"title"`       // Required!
	Description   string         `xml:"description"` // Required!
	PubDate       string         `xml:"pubDate,omitempty"`
	LastBuildDate string         `xml:"lastBuildDate,omitempty"`
	SelfLink      *atom.AtomLink `xml:"atom:link"`
	Link          string         `xml:"link"`                 // Required!
	Creator       string         `xml:"dc:creator,omitempty"` // Author
	Items         []*RssItem     `xml:"item"`
}

func NewRssFeed(f *models.Feed) *RssContainer {
	feed := &RssFeed{
		Generator:     rssGenerator,
		Language:      defaultLanguage,
		Title:         f.Title,
		Description:   f.Description,
		PubDate:       f.PublishedString(time.RFC1123Z),
		LastBuildDate: f.UpdatedString(time.RFC1123Z),
		SelfLink:      atom.NewSelfLink(f.SelfLink, RssMime),
		Link:          f.Link,
		Creator:       f.Author,
	}

	if len(f.Language) > 0 {
		feed.Language = f.Language
	}

	itemCount := len(f.Entries)
	if itemCount > 0 {
		items := make([]*RssItem, itemCount)
		i := 0

		for _, e := range f.Entries {
			if e != nil {
				items[i] = NewRssItem(e)
				i++
			}
		}

		if i > 0 {
			feed.Items = items[0:i]
		}
	}

	return wrapFeed(feed)
}
