package next_preflight

import (
	"strings"
	"testing"

	"github.com/imcrazytwkr/feedhub/utils/jsontree"
)

func TestParseEmpty(t *testing.T) {
	s, err := Parse(nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(s.Rows) != 0 || len(s.Order) != 0 {
		t.Fatal("empty buffer should yield no rows")
	}
}

func TestParseTextRow(t *testing.T) {
	s, err := Parse([]byte("a:T5,hello"))
	if err != nil {
		t.Fatal(err)
	}

	row := s.Rows[0xa]
	if row == nil || row.Tag != 'T' || row.Text != "hello" {
		t.Fatalf("T row = %+v", row)
	}
}

func TestParseInvalidJSONModel(t *testing.T) {
	s, err := Parse([]byte("0:{not json"))
	if err != nil {
		t.Fatal(err)
	}

	row := s.Rows[0]
	if row == nil || row.Value != nil {
		t.Fatalf("invalid JSON model: %+v", row)
	}
}

func TestParseRows(t *testing.T) {
	buf := []byte(strings.Join([]string{
		`1:I["./Client",["chunk"],"default"]`,
		`2:{"title":"Hi"}`,
		`0:"$2"`,
	}, "\n"))

	s, err := Parse(buf)
	if err != nil {
		t.Fatal(err)
	}

	if len(s.Order) != 3 || s.Rows[1].Tag != 'I' {
		t.Fatalf("rows = %v tag1=%q", s.Order, s.Rows[1].Tag)
	}

	if s.Rows[2].Value.GetString("title") != "Hi" {
		t.Fatal("row 2")
	}

	if s.Rows[1].Value == nil {
		t.Fatal("I-tag JSON should parse")
	}
}

func TestParseJunkBeforeID(t *testing.T) {
	s, err := Parse([]byte(`x0:{"a":1}`))
	if err != nil {
		t.Fatal(err)
	}

	if s.Rows[0] == nil || s.Rows[0].Value.GetFloat64("a") != 1 {
		t.Fatalf("rows=%v", s.Rows)
	}
}

func TestParseUppercaseHexID(t *testing.T) {
	s, err := Parse([]byte(`B:{"x":1}`))
	if err != nil {
		t.Fatal(err)
	}

	if s.Rows[0xb] == nil || s.Rows[0xb].Value.GetFloat64("x") != 1 {
		t.Fatal("uppercase id")
	}
}

func TestParseTextLengthFallback(t *testing.T) {
	s, err := Parse([]byte("1:Tz,hi\n"))
	if err != nil {
		t.Fatal(err)
	}

	row := s.Rows[1]
	if row == nil || row.Tag != 'T' || row.Text != "z,hi" {
		t.Fatalf("T fallback = %+v", row)
	}
}

func TestParseTextThenNextRow(t *testing.T) {
	s, err := Parse([]byte("a:T5,hello\n0:{}\n"))
	if err != nil {
		t.Fatal(err)
	}

	if s.Rows[0xa] == nil || s.Rows[0xa].Text != "hello" {
		t.Fatal("T row")
	}

	if s.Rows[0] == nil || s.Rows[0].Value.Type() != jsontree.TypeObject {
		t.Fatalf("next row = %+v", s.Rows[0])
	}
}

func TestParseNonJSONTag(t *testing.T) {
	s, err := Parse([]byte("1:Hnot-json\n"))
	if err != nil {
		t.Fatal(err)
	}

	row := s.Rows[1]
	if row == nil || row.Tag != 'H' || row.Text != "not-json" || row.Value != nil {
		t.Fatalf("H row = %+v", row)
	}
}
