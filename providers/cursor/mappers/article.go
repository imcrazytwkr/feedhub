package mappers

import (
	"net/url"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/imcrazytwkr/feedhub/utils/dom"
	"github.com/imcrazytwkr/feedhub/utils/htmlutil"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func PluckArticleAuthor(root *html.Node) string {
	node := htmlquery.FindOne(root, "//head/meta[@name='author']")
	if node == nil {
		return ""
	}

	return strings.TrimSpace(htmlquery.SelectAttr(node, "content"))
}

func PluckArticleContent(root *html.Node) (string, error) {
	main := lastMain(root)
	if main == nil {
		return "", nil
	}

	prose := htmlquery.FindOne(main, ".//*[contains(@class,'prose--blog')]")
	if prose == nil {
		return "", nil
	}

	wrapper := htmlutil.Element("div", nil)
	for prose.FirstChild != nil {
		child := prose.FirstChild
		prose.RemoveChild(child)
		wrapper.AppendChild(child)
	}

	prependOgImage(root, wrapper)

	if wrapper.FirstChild == nil {
		return "", nil
	}

	cleanArticle(wrapper)

	if wrapper.FirstChild == nil {
		return "", nil
	}

	var builder strings.Builder
	err := html.Render(&builder, wrapper)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(builder.String()), nil
}

func prependOgImage(doc, wrapper *html.Node) {
	src := resolveImageSrc(metaContent(doc, "og:image"))
	if len(src) == 0 {
		return
	}

	img := htmlutil.Element("img", []html.Attribute{
		{Key: "src", Val: src},
		{Key: "alt", Val: "Post splash image"},
	})

	if wrapper.FirstChild != nil {
		wrapper.InsertBefore(img, wrapper.FirstChild)
		return
	}

	wrapper.AppendChild(img)
}

func metaContent(root *html.Node, property string) string {
	node := htmlquery.FindOne(root, "//head/meta[@property='"+property+"']")
	if node == nil {
		return ""
	}

	return strings.TrimSpace(htmlquery.SelectAttr(node, "content"))
}

func cleanArticle(root *html.Node) {
	removeNodes(root, ".//*[contains(@class,'interactive-wrapper')]")
	removeNodes(root, ".//script")
	removeNodes(root, ".//style")
	removeNodes(root, ".//button")
	removeNodes(root, ".//svg")
	removeNodes(root, ".//a[contains(@class,'anchor-link')]")

	for _, node := range htmlquery.Find(root, ".//img") {
		cleanImage(node)
	}

	for _, node := range htmlquery.Find(root, ".//a") {
		cleanLink(node)
	}

	for _, node := range htmlquery.Find(root, ".//*[self::h1 or self::h2 or self::h3 or self::h4 or self::h5 or self::h6]") {
		node.Attr = nil
	}

	for _, node := range htmlquery.Find(root, ".//figure") {
		if htmlquery.FindOne(node, ".//img") == nil && node.Parent != nil {
			node.Parent.RemoveChild(node)
		}
	}

	stripExtraAttrs(root)
}

func removeNodes(root *html.Node, query string) {
	for _, node := range htmlquery.Find(root, query) {
		if node.Parent != nil {
			node.Parent.RemoveChild(node)
		}
	}
}

func cleanImage(node *html.Node) {
	if node.Parent == nil {
		return
	}

	attrs := dom.ParseAttributes(node.Attr)
	src := attrs["src"]
	if len(src) == 0 {
		src = attrs["data-src"]
	}

	src = resolveImageSrc(src)
	if len(src) == 0 {
		node.Parent.RemoveChild(node)
		return
	}

	cleaned := map[string]string{"src": src}
	if alt := strings.TrimSpace(attrs["alt"]); len(alt) > 0 {
		cleaned["alt"] = alt
	} else {
		cleaned["alt"] = "unspecified image"
	}
	if width := attrs["width"]; len(width) > 0 {
		cleaned["width"] = width
	}
	if height := attrs["height"]; len(height) > 0 {
		cleaned["height"] = height
	}

	node.Attr = dom.SerializeAttributes(cleaned)
}

func resolveImageSrc(src string) string {
	src = strings.TrimSpace(src)
	if len(src) == 0 {
		return ""
	}

	parsed, err := url.Parse(src)
	if err != nil {
		return ""
	}

	if strings.Contains(parsed.Path, "/_next/image") {
		inner := strings.TrimSpace(parsed.Query().Get("url"))
		if len(inner) > 0 {
			return inner
		}
	}

	return absolutize(src)
}

func cleanLink(node *html.Node) {
	if node.Parent == nil {
		return
	}

	href := strings.TrimSpace(htmlquery.SelectAttr(node, "href"))
	if len(href) == 0 || href == "#" || strings.HasPrefix(href, "javascript:") {
		unwrap(node)
		return
	}

	href = absolutize(href)
	if len(href) == 0 {
		unwrap(node)
		return
	}

	node.Attr = []html.Attribute{
		{Key: "href", Val: href},
		{Key: "target", Val: "_blank"},
		{Key: "rel", Val: "noopener, nofollow"},
	}
}

func unwrap(node *html.Node) {
	parent := node.Parent
	if parent == nil {
		return
	}

	for node.FirstChild != nil {
		child := node.FirstChild
		node.RemoveChild(child)
		parent.InsertBefore(child, node)
	}

	parent.RemoveChild(node)
}

func stripExtraAttrs(root *html.Node) {
	for _, node := range htmlquery.Find(root, ".//*") {
		switch node.DataAtom {
		case atom.Img, atom.A:
			continue
		}

		node.Attr = nil
	}
}
