package atom

import (
	"encoding/xml"
	"time"

	"github.com/imcrazytwkr/feedhub/models"
)

const AtomNs = "http://www.w3.org/2005/Atom"
const AtomMime = "application/atom+xml"
const atomGenerator = "FeedHub/Atom"
const defaultLanguage = "en"

type AtomFeed struct {
	XMLName   xml.Name     `xml:"feed"`
	Xmlns     string       `xml:"xmlns,attr"`
	Generator string       `xml:"generator"`
	Language  string       `xml:"xml:lang,attr"`
	Id        string       `xml:"id"`    // Required!
	Title     string       `xml:"title"` // Required!
	Subtitle  string       `xml:"subtitle,omitempty"`
	Published string       `xml:"published,omitempty"`
	Updated   string       `xml:"updated"` // Required!
	Links     []*AtomLink  `xml:"link"`    // Self link + normal link
	Author    *AtomAuthor  `xml:"author,omitempty"`
	Entries   []*AtomEntry `xml:"entry"`
}

func NewAtomFeed(f *models.Feed) *AtomFeed {
	feed := &AtomFeed{
		Xmlns:     AtomNs,
		Generator: atomGenerator,
		Language:  defaultLanguage,
		Id:        f.Link,
		Title:     f.Title,
		Subtitle:  f.Description,
		Published: f.PublishedString(time.RFC3339),
		Updated:   f.UpdatedString(time.RFC3339),
		Author:    NewAtomAuthor(f.Author),
	}

	if len(f.Language) > 0 {
		feed.Language = f.Language
	}

	feed.Links = []*AtomLink{
		NewSelfLink(f.SelfLink, AtomMime),
		NewLink(f.Link),
	}

	entryCount := len(f.Entries)
	if entryCount > 0 {
		entries := make([]*AtomEntry, entryCount)
		i := 0

		for _, e := range f.Entries {
			if e != nil {
				entries[i] = NewAtomEntry(e)
				i++
			}
		}

		if i > 0 {
			feed.Entries = entries[0:i]
		}
	}

	return feed
}
