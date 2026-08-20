package portable_text

import "github.com/imcrazytwkr/feedhub/utils/jsontree"

func isSpan(v *jsontree.Value) bool {
	if v.GetString("_type") != "span" || v.Get("text").Type() != jsontree.TypeString {
		return false
	}

	marks := v.Get("marks")
	if marks.Type() == jsontree.TypeNull {
		return true
	}

	if marks.Type() != jsontree.TypeArray {
		return false
	}

	for _, m := range v.GetArray("marks") {
		if m.Type() != jsontree.TypeString {
			return false
		}
	}

	return true
}

func decodeStrings(items []*jsontree.Value) []string {
	if len(items) == 0 {
		return nil
	}

	out := make([]string, len(items))
	for i, el := range items {
		out[i] = el.GetString()
	}

	return out
}
