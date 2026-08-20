package jsontree

import (
	"encoding/json"
	"math"
	"testing"
)

func TestParsePrimitives(t *testing.T) {
	cases := []struct {
		in   string
		typ  Type
		str  string
		num  float64
		bool bool
	}{
		{`null`, TypeNull, "", 0, false},
		{`true`, TypeTrue, "", 0, true},
		{`false`, TypeFalse, "", 0, false},
		{`"Hi"`, TypeString, "Hi", 0, false},
		{`1`, TypeNumber, "", 1, false},
		{`1.5`, TypeNumber, "", 1.5, false},
		{`  "x"  `, TypeString, "x", 0, false},
		{`  null`, TypeNull, "", 0, false},
		{" \ttrue", TypeTrue, "", 0, true},
		{"\nfalse", TypeFalse, "", 0, false},
	}

	for _, tc := range cases {
		v, err := Parse([]byte(tc.in))
		if err != nil {
			t.Fatalf("Parse(%s): %v", tc.in, err)
		}
		if v.Type() != tc.typ {
			t.Fatalf("Parse(%s) type = %v, want %v", tc.in, v.Type(), tc.typ)
		}
		if v.GetString() != tc.str {
			t.Fatalf("Parse(%s) string = %q", tc.in, v.GetString())
		}
		if v.GetFloat64() != tc.num {
			t.Fatalf("Parse(%s) number = %v", tc.in, v.GetFloat64())
		}
		if v.GetBool() != tc.bool {
			t.Fatalf("Parse(%s) bool = %v", tc.in, v.GetBool())
		}
	}
}

func TestParseRejectsInvalid(t *testing.T) {
	for _, in := range []string{
		"", "   ", "{", `[1,]`, `{'a':1}`, `nul`, `1 2`,
		`tru`, `fals`, `"unterminated`, `[1,{]`, `{"a": tru}`,
	} {
		_, err := Parse([]byte(in))
		if err == nil {
			t.Fatalf("Parse(%q) succeeded", in)
		}
	}
}

func TestParseEscapedString(t *testing.T) {
	v, err := Parse([]byte(`"a\"b\nc"`))
	if err != nil {
		t.Fatal(err)
	}

	if v.GetString() != "a\"b\nc" {
		t.Fatalf("got %q", v.GetString())
	}

	v2, err := Parse(v.MarshalTo(nil))
	if err != nil {
		t.Fatal(err)
	}

	if v2.GetString() != v.GetString() {
		t.Fatalf("%q vs %q", v2.GetString(), v.GetString())
	}
}

func TestUnmarshalJSONReuse(t *testing.T) {
	var v Value
	if err := json.Unmarshal([]byte(`{"a":1}`), &v); err != nil {
		t.Fatal(err)
	}

	if err := json.Unmarshal([]byte(`"x"`), &v); err != nil {
		t.Fatal(err)
	}

	if v.Type() != TypeString || v.GetString() != "x" || v.GetObject() != nil {
		t.Fatalf("reuse leftover: type=%v obj=%v", v.Type(), v.GetObject())
	}
}

func TestParseNaN(t *testing.T) {
	for _, in := range []string{`NaN`, `nan`, `NAN`, `nAn`, `naN`} {
		v, err := Parse([]byte(in))
		if err != nil {
			t.Fatalf("Parse(%s): %v", in, err)
		}
		if v.Type() != TypeNumber || !math.IsNaN(v.GetFloat64()) {
			t.Fatalf("Parse(%s) type=%v num=%v", in, v.Type(), v.GetFloat64())
		}
	}
}
