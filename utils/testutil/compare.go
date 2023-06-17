package testutil

import "reflect"

func SlicesEqual[V comparable](a []V, b []V) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if !reflect.DeepEqual(v, b[i]) {
			return false
		}
	}

	return true
}
