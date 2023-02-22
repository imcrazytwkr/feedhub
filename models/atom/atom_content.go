package atom

type AtomContent struct {
	Content string `xml:",chardata"`
	Type    string `xml:"type,attr"`
}

const htmlType = "html"

func NewAtomContent(text string) *AtomContent {
	if len(text) == 0 {
		return nil
	}

	return &AtomContent{
		Content: text,
		Type:    htmlType,
	}
}
