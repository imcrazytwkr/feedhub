package portable_text

import (
	"github.com/imcrazytwkr/feedhub/utils/htmlutil"
	"github.com/imcrazytwkr/feedhub/utils/jsontree"
	"golang.org/x/net/html"
)

func ToHTML(v *jsontree.Value) *html.Node {
	nodes := decodeRoot(v)
	if len(nodes) == 0 {
		return nil
	}

	root := htmlutil.Element("div", nil)
	for _, n := range nestLists(nodes) {
		render(root, n, false)
	}

	if root.FirstChild == nil {
		return nil
	}

	return root
}
