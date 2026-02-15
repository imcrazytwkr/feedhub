package feedutil

import (
	"errors"
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/imcrazytwkr/feedhub/utils/xmlutil"
)

// @NOTE: this resides in feedutil because this sanitization
// implicitly renders malformed or improperly escaped HTML
// within feed entry content in a valid form
func SanitizeContent(content *xmlquery.Node) (string, error) {
	if content == nil {
		return "", nil
	}

	err := assertType(content)
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	for child := content.FirstChild; child != nil; child = child.NextSibling {
		err = xmlutil.RenderNode(&sb, child)
		if err != nil {
			return "", err
		}
	}

	return strings.TrimSpace(sb.String()), nil
}

func assertType(node *xmlquery.Node) error {
	switch node.Type {
	case xmlquery.DocumentNode, xmlquery.DeclarationNode:
		return errors.New("invalid structure: DocumentNode not at root of XML tree")
	case xmlquery.AttributeNode:
		return fmt.Errorf("invalid structure: AttrubuteNode is a child of (%s) node", xmlutil.NodeTypeAsString(node.Parent.Type))
	case xmlquery.ElementNode, xmlquery.TextNode, xmlquery.CharDataNode, xmlquery.CommentNode:
		return nil
	default:
		return fmt.Errorf("invalid node type: %d", node.Type)
	}
}
