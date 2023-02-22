package models

import (
	"time"

	"github.com/imcrazytwkr/feedhub/utils/timeutil"
)

type Feed struct {
	Title       string
	Description string
	Language    string
	Published   time.Time
	Updated     time.Time
	SelfLink    string
	Link        string
	Author      string
	Entries     []*Entry
}

func (f *Feed) PublishedString(format string) string {
	return timeutil.FormatAnyOfTwo(format, f.Published, f.Updated)
}

func (f *Feed) UpdatedString(format string) string {
	return timeutil.FormatMaxOfTwo(format, f.Updated, f.Published)
}
