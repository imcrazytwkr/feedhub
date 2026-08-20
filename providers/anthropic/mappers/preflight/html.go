package preflight

import (
	"github.com/imcrazytwkr/feedhub/utils/htmlutil"
	"golang.org/x/net/html"
)

func wrap(tag string, attrs []html.Attribute, kid *html.Node) *html.Node {
	n := htmlutil.Element(tag, attrs)
	n.AppendChild(kid)
	return n
}
