package jsontree

func StringValue(s string) *Value {
	return &Value{t: TypeString, s: s}
}

func ArrayValue(items []*Value) *Value {
	return &Value{t: TypeArray, a: items}
}

func ObjectValue(entries map[string]*Value) *Value {
	if len(entries) == 0 {
		entries = map[string]*Value{}
	}

	return &Value{t: TypeObject, o: Object{kvs: entries}}
}
