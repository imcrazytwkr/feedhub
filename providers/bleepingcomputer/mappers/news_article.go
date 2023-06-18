package mappers

import (
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/imcrazytwkr/feedhub/utils/dom"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func PickArticleDescription(root *html.Node) string {
	description := htmlquery.FindOne(root, "//head/meta[@name='description']")
	if description == nil {
		return ""
	}

	return strings.TrimSpace(htmlquery.SelectAttr(description, "content"))
}

func PickArticleContents(root *html.Node) (string, error) {
	article := htmlquery.FindOne(root, "//div[@class='articleBody']")
	if article == nil || article.FirstChild == nil {
		return "", nil
	}

	// @TODO: handle noscript unwrapping

	// Ad-hoc trash removal
	for _, node := range htmlquery.Find(article, "//div[contains(@class,'related-article')]") {
		node.Parent.RemoveChild(node)
	}

	for _, node := range htmlquery.Find(article, "//img") {
		cleanImage(node)
	}

	for _, node := range htmlquery.Find(article, "//a") {
		cleanLink(node)
	}

	for _, node := range htmlquery.Find(article, "//p") {
		cleanParagraph(node)
	}

	for _, node := range htmlquery.Find(article, "//*[self::s or self::strike]") {
		fixStrikethrough(node)
	}

	// Cleaning up orphaned source tags
	for _, node := range htmlquery.Find(article, "//*[local-name(.) !='audio']/source") {
		node.Parent.RemoveChild(node)
	}

	builder := strings.Builder{}

	// This skips root element
	for child := article.FirstChild; child != nil; child = child.NextSibling {
		err := html.Render(&builder, child)
		if err != nil {
			return "", err
		}
	}

	return strings.TrimSpace(builder.String()), nil
}

func cleanImage(node *html.Node) {
	attrs := dom.ParseAttributes(node.Attr)

	src := attrs["src"]
	if len(src) == 0 {
		src = attrs["data-src"]
	}

	if len(src) == 0 {
		node.Parent.RemoveChild(node)
		return
	}

	for key, val := range attrs {
		switch key {
		case "src", "alt", "width", "height":
			if len(val) == 0 {
				delete(attrs, key)
			}
		default:
			delete(attrs, key)
		}
	}

	_, ok := attrs["alt"]
	if !ok {
		// Screen reader support
		attrs["alt"] = "unspecified image"
	}

	node.Attr = dom.SerializeAttributes(attrs)
}

func cleanLink(node *html.Node) {
	href := strings.TrimSpace(htmlquery.SelectAttr(node, "href"))
	if len(href) == 0 || href == "#" || strings.HasPrefix(href, "javascript:") {
		// Unwrapping dummy links
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			node.Parent.InsertBefore(child, node)
		}

		node.Parent.RemoveChild(node)
		return
	}

	node.Attr = []html.Attribute{
		{Key: "href", Val: href},
		{Key: "target", Val: "_blank"},
		{Key: "rel", Val: "noopener, nofollow"},
	}
}

func cleanParagraph(node *html.Node) {
	attrs := dom.ParseAttributes(node.Attr)

	className, ok := attrs["class"]
	if ok && strings.Contains(className, "bc_quote") {
		node.DataAtom = atom.Blockquote
		node.Data = node.DataAtom.String()
	}

	style, ok := attrs["style"]
	if !ok {
		node.Attr = nil
		return
	}

	// Only style attribute is allowed in paragraphs, if any
	node.Attr = []html.Attribute{
		{Key: "style", Val: strings.ToLower(style)},
	}
}

func fixStrikethrough(node *html.Node) {
	node.DataAtom = atom.Span
	node.Data = node.DataAtom.String()
	node.Attr = []html.Attribute{
		{Key: "style", Val: "text-decoration: strikethrough"},
	}
}
