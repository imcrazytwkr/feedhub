package htmlutil

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func Element(tag string, attrs []html.Attribute) *html.Node {
	return &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Lookup([]byte(tag)),
		Data:     tag,
		Attr:     attrs,
	}
}

func Text(data string) *html.Node {
	return &html.Node{
		Type: html.TextNode,
		Data: data,
	}
}
