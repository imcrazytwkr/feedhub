package next_preflight

import (
	"strings"
	"testing"

	"github.com/imcrazytwkr/feedhub/utils/jsontree"
)

func TestResolve(t *testing.T) {
	buf := []byte(strings.Join([]string{
		`1:I["./Client",["chunk"],"default"]`,
		`2:{"secret":"$undefined","title":"Hi"}`,
		`3:"$2"`,
		`0:["$","div",null,{"className":"x","children":[["$","p",null,{"children":"Hello"}],["$","$L1",null,{"children":"from-client"}]]}]`,
	}, "\n"))

	s, err := Parse(buf)
	if err != nil {
		t.Fatal(err)
	}

	meta := s.Resolve(2)
	if meta.Type() != jsontree.TypeObject {
		t.Fatalf("row 2: %v", meta.Type())
	}
	if meta.Get("secret") != nil {
		t.Fatal("$undefined should be omitted")
	}
	if meta.GetString("title") != "Hi" {
		t.Fatalf("title = %q", meta.GetString("title"))
	}

	ref := s.Resolve(3)
	if ref.GetString("title") != "Hi" {
		t.Fatalf("row 3 should resolve to row 2, got %s", ref.MarshalTo(nil))
	}

	root := s.Resolve(0)
	arr := root.GetArray()
	if len(arr) != 4 || arr[0].GetString() != "$" || arr[1].GetString() != "div" {
		t.Fatalf("host tuple = %s", root.MarshalTo(nil))
	}
}

func TestResolveText(t *testing.T) {
	s, err := Parse([]byte("1:T5,hello"))
	if err != nil {
		t.Fatal(err)
	}
	if s.Resolve(1).GetString() != "hello" {
		t.Fatalf("T row = %s", s.Resolve(1).MarshalTo(nil))
	}
}

func TestResolveCycle(t *testing.T) {
	s, err := Parse([]byte(`0:"$0"`))
	if err != nil {
		t.Fatal(err)
	}
	if s.Resolve(0) != nil {
		t.Fatal("cycle should be nil")
	}
}

func TestResolveEscape(t *testing.T) {
	s, err := Parse([]byte(`0:"$$undefined"`))
	if err != nil {
		t.Fatal(err)
	}
	if s.Resolve(0).GetString() != "$undefined" {
		t.Fatalf("$$ unescape = %q", s.Resolve(0).GetString())
	}
}

func TestResolveNilParser(t *testing.T) {
	var s *StreamParser
	if s.Resolve(0) != nil {
		t.Fatal("nil parser")
	}
}

func TestResolveModuleTag(t *testing.T) {
	s, err := Parse([]byte(strings.Join([]string{
		`1:I["./Client"]`,
		`0:"$L1"`,
	}, "\n")))
	if err != nil {
		t.Fatal(err)
	}
	if s.Resolve(0).GetString() != "$L1" {
		t.Fatalf("$L on I-tag should stay, got %s", s.Resolve(0).MarshalTo(nil))
	}
}

func TestResolveMissingAndNonModel(t *testing.T) {
	s, err := Parse([]byte(strings.Join([]string{
		`1:I["./Client"]`,
		`0:{not json`,
	}, "\n")))
	if err != nil {
		t.Fatal(err)
	}
	if s.Resolve(9) != nil {
		t.Fatal("missing row")
	}
	if s.Resolve(1) != nil {
		t.Fatal("I-tag")
	}
	if s.Resolve(0) != nil {
		t.Fatal("invalid JSON model")
	}
}

func TestResolveStringRefs(t *testing.T) {
	s, err := Parse([]byte(strings.Join([]string{
		`2:{"title":"Hi"}`,
		`0:"$"`,
		`1:"$G"`,
		`3:"$99"`,
		`4:"$@2"`,
	}, "\n")))
	if err != nil {
		t.Fatal(err)
	}
	if s.Resolve(0).GetString() != "$" {
		t.Fatal("$")
	}
	if s.Resolve(1).GetString() != "$G" {
		t.Fatal("$G")
	}
	if s.Resolve(3).GetString() != "$99" {
		t.Fatal("$99")
	}
	if s.Resolve(4).GetString("title") != "Hi" {
		t.Fatalf("$@2 = %s", s.Resolve(4).MarshalTo(nil))
	}
}

func TestResolveUnchangedContainers(t *testing.T) {
	s, err := Parse([]byte(strings.Join([]string{
		`0:[]`,
		`1:{}`,
		`2:[1,true]`,
		`3:{"a":1}`,
	}, "\n")))
	if err != nil {
		t.Fatal(err)
	}
	if s.Resolve(0) != s.Rows[0].Value {
		t.Fatal("empty array should reuse")
	}
	if s.Resolve(1) != s.Rows[1].Value {
		t.Fatal("empty object should reuse")
	}
	if s.Resolve(2) != s.Rows[2].Value {
		t.Fatal("plain array should reuse")
	}
	if s.Resolve(3) != s.Rows[3].Value {
		t.Fatal("plain object should reuse")
	}
}

func TestResolveChangedContainers(t *testing.T) {
	s, err := Parse([]byte(strings.Join([]string{
		`2:{"title":"Hi"}`,
		`0:["$2"]`,
		`1:{"post":"$2"}`,
	}, "\n")))
	if err != nil {
		t.Fatal(err)
	}
	arr := s.Resolve(0)
	if arr == s.Rows[0].Value || arr.GetString("0", "title") != "Hi" {
		t.Fatalf("array ref = %s", arr.MarshalTo(nil))
	}
	obj := s.Resolve(1)
	if obj == s.Rows[1].Value || obj.GetString("post", "title") != "Hi" {
		t.Fatalf("object ref = %s", obj.MarshalTo(nil))
	}
}

func TestResolveNilArraySlot(t *testing.T) {
	s := &StreamParser{Rows: map[int]*Row{
		0: {ID: 0, Value: jsontree.ArrayValue([]*jsontree.Value{nil, jsontree.StringValue("x")})},
	}}
	got := s.Resolve(0)
	arr := got.GetArray()
	if len(arr) != 2 || arr[0] != nil || arr[1].GetString() != "x" {
		t.Fatalf("nil slot = %s", got.MarshalTo(nil))
	}
}
