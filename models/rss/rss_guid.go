package rss

import (
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/utils/feedutil"
)

type RssItemGuid struct {
	Guid        string `xml:",chardata"`
	IsPermaLink bool   `xml:"isPermaLink,attr"`
}

func NewRssItemGuid(entry *models.Entry) *RssItemGuid {
	uuid := feedutil.GenerateId(entry)
	if len(uuid) == 0 {
		return nil
	}

	return &RssItemGuid{
		Guid:        uuid,
		IsPermaLink: uuid == entry.Link,
	}
}
