package xmlutil

import (
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/imcrazytwkr/feedhub/utils/htmlutil"
	"github.com/imcrazytwkr/feedhub/utils/stringutil"
)

func RenderNode(sb *strings.Builder, node *xmlquery.Node) error {
	switch node.Type {
	case xmlquery.TextNode, xmlquery.CharDataNode:
		return renderText(sb, node)
	case xmlquery.ElementNode:
		return renderElement(sb, node)
	case xmlquery.CommentNode:
		return nil
	default:
		return fmt.Errorf("invalid node type: %s", NodeTypeAsString(node.Type))
	}
}

func renderText(sb *strings.Builder, node *xmlquery.Node) (err error) {
	data := stringutil.TrimTrailingSpace(node.Data)
	if len(data) > 0 {
		_, err = sb.WriteString(data)
		if err != nil {
			return
		}
	}

	// Important to preserve pre-code formatting
	err = sb.WriteByte('\n')
	return
}

// @NOTE: this code has been extracted for the sake of readability and should not
// be called separately
func renderElement(sb *strings.Builder, node *xmlquery.Node) (err error) {
	tag := strings.ToLower(node.Data)

	err = sb.WriteByte('<')
	if err != nil {
		return
	}

	_, err = sb.WriteString(tag)
	if err != nil {
		return
	}

	for _, attr := range node.Attr {
		err = RenderAttr(sb, &attr)
		if err != nil {
			return
		}
	}

	if htmlutil.IsVoidElement(tag) {
		if node.FirstChild != nil {
			return fmt.Errorf("void element <%s> has child nodes", tag)
		}

		_, err = sb.WriteString("/>")
		return
	}

	err = sb.WriteByte('>')
	if err != nil {
		return
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		err = RenderNode(sb, child)
		if err != nil {
			return
		}
	}

	_, err = sb.WriteString("</")
	if err != nil {
		return
	}

	_, err = sb.WriteString(tag)
	if err != nil {
		return
	}

	err = sb.WriteByte('>')
	return
}

func RenderAttr(sb *strings.Builder, attr *xmlquery.Attr) (err error) {
	err = sb.WriteByte(' ')
	if err != nil {
		return
	}

	name := attr.Name
	if len(name.Space) > 0 {
		_, err = sb.WriteString(attr.NamespaceURI)
		if err != nil {
			return
		}

		err = sb.WriteByte(':')
		if err != nil {
			return
		}
	}

	_, err = sb.WriteString(attr.Name.Local)
	if err != nil {
		return
	}

	_, err = sb.WriteString(`="`)
	if err != nil {
		return
	}

	err = htmlutil.WriteEscaped(sb, attr.Value)
	if err != nil {
		return
	}

	err = sb.WriteByte('"')
	return
}
