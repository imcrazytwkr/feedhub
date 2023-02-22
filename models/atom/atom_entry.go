package atom

import (
	"time"

	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/utils/feedutil"
)

type AtomEntry struct {
	Id         string       `xml:"id"`    // Required!
	Title      string       `xml:"title"` // Required!
	Published  string       `xml:"published,omitempty"`
	Updated    string       `xml:"updated"`          // Required!
	Author     *AtomAuthor  `xml:"author,omitempty"` // Required if feed has no author
	Link       *AtomLink    `xml:"link,omitempty"`
	Summary    *AtomContent `xml:"summary,omitempty"`
	Content    *AtomContent `xml:"content"` // Required!
	Categories []string     `xml:"category,omitempty"`
}

func NewAtomEntry(e *models.Entry) *AtomEntry {
	return &AtomEntry{
		Id:        feedutil.GenerateId(e),
		Title:     e.Title,
		Published: e.PublishedString(time.RFC3339),
		Updated:   e.UpdatedString(time.RFC3339),
		Author:    NewAtomAuthor(e.Author),
		Link:      NewLink(e.Link),
		Summary:   NewAtomContent(e.Description),
		Content:   NewAtomContent(e.Content),
	}
}
