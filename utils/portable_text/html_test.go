package portable_text_test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/imcrazytwkr/feedhub/utils/jsontree"
	pt "github.com/imcrazytwkr/feedhub/utils/portable_text"
	"github.com/imcrazytwkr/feedhub/utils/testutil"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func TestToHTML(t *testing.T) {
	matches, err := filepath.Glob("testdata/*.json")
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) == 0 {
		t.Fatal("no testdata/*.json")
	}

	for _, path := range matches {
		name := strings.TrimSuffix(filepath.Base(path), ".json")
		t.Run(name, func(t *testing.T) {
			src, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			v, err := jsontree.Parse(src)
			if err != nil {
				t.Fatal(err)
			}

			got := pt.ToHTML(v)
			wantPath := filepath.Join("testdata", name+".html")
			wantSrc, err := os.ReadFile(wantPath)
			if err != nil {
				if os.IsNotExist(err) {
					if got != nil {
						t.Fatalf("ToHTML = %s, want nil", renderHTML(t, got))
					}
					return
				}
				t.Fatal(err)
			}

			want := parseWrapped(t, strings.TrimSpace(string(wantSrc)))
			if got == nil {
				t.Fatalf("ToHTML = nil, want %s", strings.TrimSpace(string(wantSrc)))
			}
			if !testutil.DomsEqual(got, want) {
				t.Fatalf("got %s\nwant %s", renderHTML(t, got), renderHTML(t, want))
			}
		})
	}
}

func TestToHTMLNil(t *testing.T) {
	if pt.ToHTML(nil) != nil {
		t.Fatal("nil input")
	}
}

func parseWrapped(t *testing.T, inner string) *html.Node {
	t.Helper()
	nodes, err := html.ParseFragment(strings.NewReader("<div>"+inner+"</div>"), &html.Node{
		Type:     html.ElementNode,
		Data:     "body",
		DataAtom: atom.Body,
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(nodes) != 1 {
		t.Fatalf("expected one root, got %d", len(nodes))
	}
	return nodes[0]
}

func renderHTML(t *testing.T, n *html.Node) string {
	t.Helper()
	var buf bytes.Buffer
	if err := html.Render(&buf, n); err != nil {
		t.Fatal(err)
	}
	return buf.String()
}
