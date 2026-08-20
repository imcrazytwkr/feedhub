package portable_text

import "github.com/imcrazytwkr/feedhub/utils/jsontree"

func isBlock(v *jsontree.Value) bool {
	typ := v.GetString("_type")
	if typ == "" || typ[0] == '@' {
		return false
	}

	if v.Get("children").Type() != jsontree.TypeArray {
		return false
	}

	for _, child := range v.GetArray("children") {
		if child.GetObject() == nil || len(child.GetString("_type")) == 0 {
			return false
		}
	}

	defs := v.Get("markDefs")
	if defs.Type() != jsontree.TypeNull && defs.Type() != jsontree.TypeArray {
		return false
	}

	for _, def := range v.GetArray("markDefs") {
		if def.GetString("_key") == "" {
			return false
		}
	}

	return true
}

func decodeBlock(v *jsontree.Value) *node {
	n := &node{
		kind:     kindBlock,
		style:    v.GetString("style"),
		listItem: v.GetString("listItem"),
	}

	if v.Get("level").Type() == jsontree.TypeNumber {
		n.level = int(v.GetFloat64("level"))
	}

	for _, child := range v.GetArray("children") {
		n.children = append(n.children, decodeChild(child))
	}

	if defs := v.GetArray("markDefs"); len(defs) > 0 {
		n.defs = make(map[string]*markDef, len(defs))
		for _, def := range defs {
			key := def.GetString("_key")
			n.defs[key] = &markDef{
				typ:  def.GetString("_type"),
				href: def.GetString("href"),
			}
		}
	}

	return n
}

func decodeChild(v *jsontree.Value) *node {
	if isSpan(v) {
		return &node{
			kind:  kindSpan,
			marks: decodeStrings(v.GetArray("marks")),
			text:  v.GetString("text"),
		}
	}

	return &node{kind: kindUnknown, typ: v.GetString("_type")}
}
