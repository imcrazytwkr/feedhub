package rss

import (
	"encoding/xml"

	"github.com/imcrazytwkr/feedhub/models/atom"
)

const RssVersion = "2.0"
const ContentNs = "http://purl.org/rss/1.0/modules/content/"
const DublinCoreNs = "http://purl.org/dc/elements/1.1/"

type RssContainer struct {
	XMLName             xml.Name `xml:"rss"`
	Version             string   `xml:"version,attr"`
	AtomNamespace       string   `xml:"xmlns:atom,attr"`
	ContentNamespace    string   `xml:"xmlns:content,attr"`
	DublinCoreNamespace string   `xml:"xmlns:dc,attr"`
	Channel             *RssFeed
}

func wrapFeed(feed *RssFeed) *RssContainer {
	return &RssContainer{
		Version:             RssVersion,
		AtomNamespace:       atom.AtomNs,
		ContentNamespace:    ContentNs,
		DublinCoreNamespace: DublinCoreNs,
		Channel:             feed,
	}
}
