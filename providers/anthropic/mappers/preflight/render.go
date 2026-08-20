package preflight

import (
	"strings"

	"github.com/imcrazytwkr/feedhub/utils/htmlutil"
	pt "github.com/imcrazytwkr/feedhub/utils/portable_text"
	"golang.org/x/net/html"
)

func RenderArticleHTML(preflight []byte) *html.Node {
	article := findPreflightArticle(preflight)
	if !article.CanRender() {
		return nil
	}

	root := pt.ToHTML(article.Body)
	if root == nil {
		return nil
	}

	notes := pt.ToHTML(article.FootnotesBody)
	if notes == nil {
		return root
	}

	title := strings.TrimSpace(article.FootnotesTitle)
	if len(title) == 0 {
		title = "Footnotes"
	}

	root.AppendChild(wrap("h4", nil, htmlutil.Text(title)))

	for notes.FirstChild != nil {
		kid := notes.FirstChild
		notes.RemoveChild(kid)
		root.AppendChild(kid)
	}

	return root
}
