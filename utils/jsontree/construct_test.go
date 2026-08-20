package jsontree

import "testing"

func TestString(t *testing.T) {
	v := StringValue("Hi")
	if v.Type() != TypeString || v.GetString() != "Hi" {
		t.Fatalf("String: type=%v val=%q", v.Type(), v.GetString())
	}
}

func TestArray(t *testing.T) {
	v := ArrayValue([]*Value{StringValue("$"), StringValue("div")})
	arr := v.GetArray()
	if v.Type() != TypeArray || len(arr) != 2 || arr[0].GetString() != "$" || arr[1].GetString() != "div" {
		t.Fatalf("Array = %s", v.MarshalTo(nil))
	}
}

func TestNewObject(t *testing.T) {
	v := ObjectValue(map[string]*Value{"title": StringValue("Hi")})
	if v.Type() != TypeObject || v.GetString("title") != "Hi" || v.GetObject().Len() != 1 {
		t.Fatalf("NewObject = %s", v.MarshalTo(nil))
	}

	empty := ObjectValue(nil)
	if empty.Type() != TypeObject || empty.GetObject().Len() != 0 {
		t.Fatal("NewObject(nil) should be empty object")
	}
}

func TestArrayValueNilSlot(t *testing.T) {
	v := ArrayValue([]*Value{nil, StringValue("x")})
	arr := v.GetArray()
	if len(arr) != 2 || arr[0] != nil || arr[1].GetString() != "x" {
		t.Fatal("nil array slot")
	}
}
