package jsontree

import "testing"

func TestNilObject(t *testing.T) {
	var o *Object
	if o.Get("a") != nil || o.Len() != 0 {
		t.Fatal("nil Object should be empty")
	}
}

func TestEmptyObjectLen(t *testing.T) {
	v, err := Parse([]byte(`{}`))
	if err != nil {
		t.Fatal(err)
	}

	obj := v.GetObject()
	if obj == nil || obj.Len() != 0 {
		t.Fatalf("empty object Len = %v", obj.Len())
	}

	nested, err := Parse([]byte(`{"post":{}}`))
	if err != nil {
		t.Fatal(err)
	}

	if nested.GetObject("post").Len() != 0 {
		t.Fatal("nested empty object")
	}
}

func TestObjectGet(t *testing.T) {
	v, err := Parse([]byte(`{"href":"https://x","n":1}`))
	if err != nil {
		t.Fatal(err)
	}

	obj := v.GetObject()
	if obj.Len() != 2 {
		t.Fatalf("Len = %d", obj.Len())
	}

	if obj.Get("href").GetString() != "https://x" {
		t.Fatalf("href = %q", obj.Get("href").GetString())
	}

	if obj.Get("missing") != nil {
		t.Fatal("missing key")
	}
}

func TestObjectEntries(t *testing.T) {
	v, err := Parse([]byte(`{"href":"https://x","n":1}`))
	if err != nil {
		t.Fatal(err)
	}

	obj := v.GetObject()
	got := map[string]string{}
	for k, val := range obj.Entries() {
		if val.Type() == TypeString {
			got[k] = val.GetString()
		}
		if val.Type() == TypeNumber {
			got[k] = "num"
		}
	}

	if got["href"] != "https://x" || got["n"] != "num" {
		t.Fatalf("Entries = %v", got)
	}

	n := 0
	for range obj.Entries() {
		n++
		break
	}

	if n != 1 {
		t.Fatalf("early break n=%d", n)
	}

	empty, err := Parse([]byte(`{}`))
	if err != nil {
		t.Fatal(err)
	}

	for range empty.GetObject().Entries() {
		t.Fatal("empty Entries should be empty")
	}

	var o *Object
	for range o.Entries() {
		t.Fatal("nil Entries should be empty")
	}
}
