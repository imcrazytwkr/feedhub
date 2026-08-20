package next_preflight

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestCollectConcatenatesChunks(t *testing.T) {
	doc := `<html><body>
<script>(self.__next_f=self.__next_f||[]).push([0])</script>
<script>self.__next_f.push([1,"0:{\"a\":"])</script>
<script>self.__next_f.push([1,"1}\n"])</script>
</body></html>`

	root, err := html.Parse(strings.NewReader(doc))
	if err != nil {
		t.Fatal(err)
	}

	buf := Collect(root)
	if string(buf) != `0:{"a":1}`+"\n" {
		t.Fatalf("collected %q", buf)
	}

	s, err := Parse(buf)
	if err != nil {
		t.Fatal(err)
	}
	if s.Resolve(0).GetFloat64("a") != 1 {
		t.Fatalf("resolved %s", s.Resolve(0).MarshalTo(nil))
	}
}
