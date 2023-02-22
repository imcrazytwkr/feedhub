package rss

type RssContent struct {
	Content string `xml:",cdata"`
	Type    string `xml:"type,attr"`
}

const htmlType = "html"

func NewRssContent(text string) *RssContent {
	if len(text) == 0 {
		return nil
	}

	return &RssContent{
		Content: text,
		Type:    htmlType,
	}
}
