package next_preflight

import "github.com/imcrazytwkr/feedhub/utils/jsontree"

const undefinedSentinel = "$undefined"

func (s *StreamParser) Resolve(id int) *jsontree.Value {
	if s == nil {
		return nil
	}

	return s.resolveRow(id, make(map[int]bool))
}

func (s *StreamParser) resolveRow(id int, stack map[int]bool) *jsontree.Value {
	row := s.Rows[id]
	if row == nil {
		return nil
	}

	if row.Tag == 'T' {
		return jsontree.StringValue(row.Text)
	}

	if row.Tag != 0 || row.Value == nil {
		return nil
	}

	if stack[id] {
		return nil
	}

	stack[id] = true
	defer delete(stack, id)

	return s.resolveValue(row.Value, stack)
}

func (s *StreamParser) resolveValue(v *jsontree.Value, stack map[int]bool) *jsontree.Value {
	if v == nil {
		return nil
	}

	switch v.Type() {
	case jsontree.TypeString:
		return s.resolveString(v, stack)
	case jsontree.TypeArray:
		return s.resolveArray(v, stack)
	case jsontree.TypeObject:
		return s.resolveObject(v, stack)
	default:
		return v
	}
}

func isUndefined(v *jsontree.Value) bool {
	return v.Type() == jsontree.TypeString && v.GetString() == undefinedSentinel
}

func (s *StreamParser) resolveString(v *jsontree.Value, stack map[int]bool) *jsontree.Value {
	if v == nil {
		return nil
	}

	val := v.GetString()
	if len(val) == 0 || val[0] != '$' {
		return v
	}

	switch val {
	case "$":
		// Self-ref
		return v
	case undefinedSentinel:
		return nil
	}

	rest := val[1:]

	if rest[0] == '$' {
		return jsontree.StringValue(rest)
	}

	kind := byte(0)
	switch rest[0] {
	case 'L', '@':
		kind = rest[0]
		rest = rest[1:]
	}

	id, n := parseHexPrefix(rest)
	if n == 0 {
		return v
	}

	row := s.Rows[id]
	if row == nil {
		return v
	}

	if kind == 'L' && row.Tag != 0 && row.Tag != 'T' {
		return v
	}

	return s.resolveRow(id, stack)
}

func (s *StreamParser) resolveArray(v *jsontree.Value, stack map[int]bool) *jsontree.Value {
	if v == nil {
		return nil
	}

	items := v.GetArray()
	if len(items) == 0 {
		return v
	}

	out := make([]*jsontree.Value, len(items))
	changed := false

	for i, el := range items {
		// Should not optimise-out undefined sentinels because
		// otherwise tuple indexing breaks
		resolved := s.resolveValue(el, stack)
		out[i] = resolved
		if resolved != el {
			changed = true
		}
	}

	if !changed {
		return v
	}

	return jsontree.ArrayValue(out)
}

func (s *StreamParser) resolveObject(v *jsontree.Value, stack map[int]bool) *jsontree.Value {
	if v == nil {
		return nil
	}

	o := v.GetObject()
	if o.Len() == 0 {
		return v
	}

	out := make(map[string]*jsontree.Value, o.Len())
	changed := false

	for k, el := range o.Entries() {
		if isUndefined(el) {
			changed = true
			continue
		}

		resolved := s.resolveValue(el, stack)
		out[k] = resolved
		if resolved != el {
			changed = true
		}
	}

	if !changed {
		return v
	}

	return jsontree.ObjectValue(out)
}
