package testutil

import (
	"github.com/imcrazytwkr/feedhub/utils/dom"
	"golang.org/x/net/html"
)

func DomsEqual(a *html.Node, b *html.Node) bool {
	if a.Type != b.Type || a.Data != b.Data {
		return false
	}

	aAttr := dom.ParseAttributes(a.Attr)
	bAttr := dom.ParseAttributes(b.Attr)

	if len(aAttr) != len(bAttr) {
		return false
	}

	for key, aVal := range aAttr {
		bVal, ok := bAttr[key]
		if !ok || aVal != bVal {
			return false
		}
	}

	aChildren := parseChildren(a)
	bChildren := parseChildren(b)

	if len(aChildren) != len(bChildren) {
		return false
	}

	if len(aChildren) == 0 {
		return true
	}

	for i, child := range aChildren {
		if !DomsEqual(child, bChildren[i]) {
			return false
		}
	}

	return true
}

// @NOTE: this function is not optimal but since it's only used in testing, it doesn't matter
func parseChildren(node *html.Node) []*html.Node {
	var result []*html.Node
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		result = append(result, child)
	}

	return result
}
