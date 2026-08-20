package portable_text

import (
	"fmt"
	"strings"

	"github.com/imcrazytwkr/feedhub/utils/htmlutil"
	"golang.org/x/net/html"
)

func render(parent *html.Node, n *node, inline bool) {
	if parent == nil || n == nil {
		return
	}

	switch n.kind {
	case kindList:
		appendHTML(parent, renderList(n))
	case kindBlock:
		if len(n.listItem) > 0 {
			appendHTML(parent, renderListItem(n))
			return
		}

		appendHTML(parent, renderBlock(n))
	case kindMark:
		renderMark(parent, n)
	case kindText:
		if child := renderText(n); child != nil {
			appendHTML(parent, child)
		}
	default:
		appendHTML(parent, renderUnknown(n, inline))
	}
}

func renderList(n *node) *html.Node {
	tag := "ul"
	if n.listItem == "number" {
		tag = "ol"
	}
	list := htmlutil.Element(tag, nil)
	for _, child := range n.children {
		if item := renderListItem(child); item != nil {
			appendHTML(list, item)
		}
	}
	return list
}

func renderListItem(n *node) *html.Node {
	if n == nil {
		return nil
	}

	li := htmlutil.Element("li", nil)
	if n.style != "" && n.style != "normal" {
		item := *n
		item.listItem = ""
		appendHTML(li, renderBlock(&item))
		return li
	}
	for _, child := range nestMarks(n) {
		render(li, child, true)
	}
	return li
}

func renderBlock(n *node) *html.Node {
	style := n.style
	if style == "" {
		style = "normal"
	}
	el := htmlutil.Element(blockTag(style), nil)
	for _, child := range nestMarks(n) {
		render(el, child, true)
	}
	return el
}

func blockTag(style string) string {
	switch style {
	case "h1", "h2", "h3", "h4", "h5", "h6", "blockquote":
		return style
	default:
		return "p"
	}
}

func renderMark(parent *html.Node, n *node) {
	switch n.typ {
	case "em", "strong", "code":
		el := htmlutil.Element(n.typ, nil)
		renderChildren(el, n)
		appendHTML(parent, el)
	case "underline":
		el := htmlutil.Element("span", []html.Attribute{{Key: "style", Val: "text-decoration:underline"}})
		renderChildren(el, n)
		appendHTML(parent, el)
	case "strike-through":
		el := htmlutil.Element("del", nil)
		renderChildren(el, n)
		appendHTML(parent, el)
	case "link":
		href := ""
		if n.def != nil {
			href = n.def.href
		}
		if !uriLooksSafe(href) {
			renderChildren(parent, n)
			return
		}
		el := htmlutil.Element("a", []html.Attribute{{Key: "href", Val: href}})
		renderChildren(el, n)
		appendHTML(parent, el)
	default:
		el := htmlutil.Element("span", []html.Attribute{{Key: "class", Val: "unknown__pt__mark__" + n.typ}})
		renderChildren(el, n)
		appendHTML(parent, el)
	}
}

func renderChildren(parent *html.Node, n *node) {
	for _, child := range n.children {
		render(parent, child, true)
	}
}

func renderText(n *node) *html.Node {
	switch n.text {
	case "\n":
		return htmlutil.Element("br", nil)
	case "":
		return nil
	default:
		return htmlutil.Text(preserveSpaces(n.text))
	}
}

func renderUnknown(n *node, inline bool) *html.Node {
	msg := fmt.Sprintf("Unknown block type \"%s\", specify a component for it in the `components.types` option", n.typ)
	tag := "div"
	if inline {
		tag = "span"
	}
	el := htmlutil.Element(tag, []html.Attribute{{Key: "style", Val: "display:none"}})
	el.AppendChild(htmlutil.Text(msg))
	return el
}

func preserveSpaces(s string) string {
	var b strings.Builder
	n := 0
	for _, r := range s {
		if r == ' ' {
			n++
			continue
		}
		writeSpaces(&b, n)
		n = 0
		b.WriteRune(r)
	}
	writeSpaces(&b, n)
	return b.String()
}

func writeSpaces(b *strings.Builder, n int) {
	if n == 0 {
		return
	}
	if n == 1 {
		b.WriteByte(' ')
		return
	}
	for i := 0; i < n-1; i++ {
		b.WriteRune('\u00a0')
	}
	b.WriteByte(' ')
}

func appendHTML(parent, child *html.Node) {
	if parent == nil || child == nil {
		return
	}
	if child.Type == html.TextNode && parent.LastChild != nil && parent.LastChild.Type == html.TextNode {
		parent.LastChild.Data += child.Data
		return
	}
	parent.AppendChild(child)
}

func uriLooksSafe(uri string) bool {
	url := strings.TrimSpace(uri)
	if url == "" {
		return true
	}

	first := url[0]
	if first == '#' || first == '/' {
		return true
	}

	colon := strings.IndexByte(url, ':')
	if colon == -1 {
		return true
	}

	proto := strings.ToLower(url[:colon])
	switch proto {
	case "http", "https", "mailto", "tel":
		return true
	}

	query := strings.IndexByte(url, '?')
	if query != -1 && colon > query {
		return true
	}

	hash := strings.IndexByte(url, '#')
	if hash != -1 && colon > hash {
		return true
	}

	return false
}
