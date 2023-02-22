package models

import (
	"time"

	"github.com/imcrazytwkr/feedhub/utils/timeutil"
)

type Entry struct {
	Id          string
	Title       string
	Published   time.Time
	Updated     time.Time
	Author      string
	Link        string
	Description string
	Content     string
}

func (e *Entry) PublishedString(format string) string {
	return timeutil.FormatAnyOfTwo(format, e.Published, e.Updated)
}

func (e *Entry) UpdatedString(format string) string {
	return timeutil.FormatMaxOfTwo(format, e.Updated, e.Published)
}
