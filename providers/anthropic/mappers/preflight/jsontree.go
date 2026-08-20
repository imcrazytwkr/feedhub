package preflight

import (
	"iter"

	"github.com/imcrazytwkr/feedhub/utils/jsontree"
)

func walkJSON(v *jsontree.Value) iter.Seq[*jsontree.Value] {
	return func(yield func(*jsontree.Value) bool) {
		var walk func(*jsontree.Value) bool
		walk = func(n *jsontree.Value) bool {
			if n == nil {
				return true
			}

			if !yield(n) {
				return false
			}

			switch n.Type() {
			case jsontree.TypeArray:
				for _, el := range n.GetArray() {
					if !walk(el) {
						return false
					}
				}
			case jsontree.TypeObject:
				for _, el := range n.GetObject().Entries() {
					if !walk(el) {
						return false
					}
				}
			}

			return true
		}
		walk(v)
	}
}
