package xmlutil

import (
	"github.com/antchfx/xmlquery"
)

func NodeTypeAsString(nodeType xmlquery.NodeType) string {
	switch nodeType {
	case xmlquery.DocumentNode:
		return "DocumentNode"
	case xmlquery.DeclarationNode:
		return "DeclarationNode"
	case xmlquery.ElementNode:
		return "ElementNode"
	case xmlquery.TextNode:
		return "TextNode"
	case xmlquery.CharDataNode:
		return "CharDataNode"
	case xmlquery.CommentNode:
		return "CharDataNode"
	case xmlquery.AttributeNode:
		return "CharDataNode"
	default:
		return "UnknownNode"
	}
}
