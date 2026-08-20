package portable_text

import "github.com/imcrazytwkr/feedhub/utils/jsontree"

func decodeRoot(v *jsontree.Value) []*node {
	if v == nil {
		return nil
	}

	switch v.Type() {
	case jsontree.TypeArray:
		return decodeNodes(v.GetArray())
	case jsontree.TypeObject:
		return []*node{decodeNode(v)}
	default:
		return nil
	}
}

func decodeNodes(items []*jsontree.Value) []*node {
	if len(items) == 0 {
		return nil
	}

	out := make([]*node, 0, len(items))
	for _, el := range items {
		if el == nil {
			continue
		}

		out = append(out, decodeNode(el))
	}

	return out
}

func decodeNode(v *jsontree.Value) *node {
	if isBlock(v) {
		return decodeBlock(v)
	}

	return &node{kind: kindUnknown, typ: v.GetString("_type")}
}
