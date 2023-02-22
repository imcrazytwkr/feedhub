package atom

const relSelf = "self"

type AtomLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr,omitempty"`
	Type string `xml:"type,attr,omitempty"`
}

func NewLink(href string) *AtomLink {
	if len(href) == 0 {
		return nil
	}

	return &AtomLink{
		Href: href,
	}
}

func NewSelfLink(href string, mime string) *AtomLink {
	if len(href) == 0 {
		return nil
	}

	return &AtomLink{
		Href: href,
		Rel:  relSelf,
		Type: mime,
	}
}
