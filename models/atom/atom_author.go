package atom

type AtomAuthor struct {
	Name  string `xml:"name"` // Required!
	Uri   string `xml:"uri,omitempty"`
	Email string `xml:"email,omitempty"`
}

func NewAtomAuthor(name string) *AtomAuthor {
	return &AtomAuthor{
		Name: name,
	}
}
