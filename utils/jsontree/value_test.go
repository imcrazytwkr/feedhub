package jsontree

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestGetPaths(t *testing.T) {
	v, err := Parse([]byte(`{"post":{"title":"Hi","n":2},"arr":["a",{"k":"v"},null]}`))
	if err != nil {
		t.Fatal(err)
	}

	if v.GetString("post", "title") != "Hi" {
		t.Fatalf("title = %q", v.GetString("post", "title"))
	}

	if v.GetFloat64("post", "n") != 2 {
		t.Fatalf("n = %v", v.GetFloat64("post", "n"))
	}

	if v.GetString("arr", "0") != "a" {
		t.Fatalf("arr.0 = %q", v.GetString("arr", "0"))
	}

	if v.GetString("arr", "1", "k") != "v" {
		t.Fatalf("arr.1.k = %q", v.GetString("arr", "1", "k"))
	}

	if got := v.Get("arr", "2"); got == nil || got.Type() != TypeNull {
		t.Fatalf("arr.2 = %v", got)
	}

	if v.Get("missing") != nil || v.Get("arr", "9") != nil || v.Get("arr", "x") != nil || v.Get("arr", "-1") != nil {
		t.Fatal("missing paths should be nil")
	}

	if v.Get("post", "n", "x") != nil {
		t.Fatal("Get into a scalar should be nil")
	}

	if v.Get() != v {
		t.Fatal("Get with no keys should return the receiver")
	}

	if v.GetString("post") != "" || v.GetFloat64("title") != 0 || v.GetBool("post") {
		t.Fatal("wrong-type Get* should be zero")
	}

	if v.GetArray("post") != nil || v.GetObject("arr") != nil || v.GetArray("post", "title") != nil {
		t.Fatal("GetArray/GetObject on the wrong container should be nil")
	}
}

func TestNilValue(t *testing.T) {
	var v *Value
	if v.Get("a") != nil || v.GetString("a") != "" || v.GetArray() != nil || v.GetObject() != nil {
		t.Fatal("nil Value Get* should be zero")
	}
	if v.Type() != TypeNull {
		t.Fatalf("nil Type = %v", v.Type())
	}
	if v.GetBool() {
		t.Fatal("nil GetBool")
	}
	if string(v.MarshalTo(nil)) != "null" {
		t.Fatalf("nil MarshalTo = %s", v.MarshalTo(nil))
	}
}

func TestEmptyArray(t *testing.T) {
	v, err := Parse([]byte(`[]`))
	if err != nil {
		t.Fatal(err)
	}
	if v.Type() != TypeArray || len(v.GetArray()) != 0 {
		t.Fatalf("empty array: type=%v len=%d", v.Type(), len(v.GetArray()))
	}
	if string(v.MarshalTo(nil)) != "[]" {
		t.Fatalf("MarshalTo = %s", v.MarshalTo(nil))
	}
}

func TestMarshalJSON(t *testing.T) {
	v, err := Parse([]byte(`["$",1,true]`))
	if err != nil {
		t.Fatal(err)
	}

	got, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, v.MarshalTo(nil)) {
		t.Fatalf("json.Marshal = %s, MarshalTo = %s", got, v.MarshalTo(nil))
	}

	v2, err := Parse(got)
	if err != nil {
		t.Fatal(err)
	}
	if v2.GetString("0") != "$" || v2.GetFloat64("1") != 1 || !v2.GetBool("2") {
		t.Fatalf("roundtrip via MarshalJSON: %s", got)
	}
}

func TestFlightHostTuple(t *testing.T) {
	v, err := Parse([]byte(`["$","div",null,{"className":"x","children":[["$","p",null,{"children":"Hello"}]]}]`))
	if err != nil {
		t.Fatal(err)
	}
	arr := v.GetArray()
	if len(arr) != 4 || arr[0].GetString() != "$" || arr[1].GetString() != "div" {
		t.Fatalf("tuple = %s", v.MarshalTo(nil))
	}
	if arr[2].Type() != TypeNull {
		t.Fatal("key should be null")
	}
	props := arr[3].GetObject()
	if props.Get("className").GetString() != "x" {
		t.Fatal("className")
	}
	kids := props.Get("children").GetArray()
	if len(kids) != 1 || kids[0].GetString("1") != "p" {
		t.Fatal("children")
	}
	if kids[0].GetString("3", "children") != "Hello" {
		t.Fatalf("text = %q", kids[0].GetString("3", "children"))
	}
}

func TestMarshalToRoundTrip(t *testing.T) {
	in := []byte(`{"title":"Hi","a":1,"ok":true,"no":false,"z":null,"arr":["$","div"],"nested":{"k":"v"}}`)
	v, err := Parse(in)
	if err != nil {
		t.Fatal(err)
	}

	v2, err := Parse(v.MarshalTo(nil))
	if err != nil {
		t.Fatal(err)
	}

	if v2.GetString("title") != "Hi" || v2.GetFloat64("a") != 1 || !v2.GetBool("ok") || v2.GetBool("no") {
		t.Fatalf("roundtrip scalars: %s", v2.MarshalTo(nil))
	}

	if v2.Get("z").Type() != TypeNull {
		t.Fatal("null")
	}

	if v2.GetString("arr", "0") != "$" || v2.GetString("nested", "k") != "v" {
		t.Fatal("nested roundtrip")
	}
}
