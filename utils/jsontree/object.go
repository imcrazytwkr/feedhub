package jsontree

import (
	"iter"
)

type Object struct {
	kvs map[string]*Value
}

func (o *Object) Get(key string) *Value {
	if o == nil {
		return nil
	}

	return o.kvs[key]
}

func (o *Object) Len() int {
	if o == nil {
		return 0
	}

	return len(o.kvs)
}

func (o *Object) Entries() iter.Seq2[string, *Value] {
	return func(yield func(string, *Value) bool) {
		if o.Len() == 0 {
			return
		}

		for k, v := range o.kvs {
			if !yield(k, v) {
				return
			}
		}
	}
}
