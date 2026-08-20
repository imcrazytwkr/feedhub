package jsontree

import "testing"

func TestTypeString(t *testing.T) {
	cases := []struct {
		typ Type
		s   string
	}{
		{TypeNull, "null"},
		{TypeObject, "object"},
		{TypeArray, "array"},
		{TypeString, "string"},
		{TypeNumber, "number"},
		{TypeTrue, "true"},
		{TypeFalse, "false"},
	}
	for _, tc := range cases {
		if got := tc.typ.String(); got != tc.s {
			t.Fatalf("%v.String() = %q, want %q", tc.typ, got, tc.s)
		}
	}
}
