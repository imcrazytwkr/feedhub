package mappers_test

import (
	"strings"
	"testing"

	"github.com/imcrazytwkr/feedhub/providers/anthropic/mappers/preflight"
	np "github.com/imcrazytwkr/feedhub/utils/next_preflight"
	"golang.org/x/net/html"
)

func TestRenderArticlePreflight(t *testing.T) {
	cases := []struct {
		source   string
		expected string
	}{
		{"testdata/news_article.html", "testdata/expected_news_preflight.html"},
		{"testdata/engineering_article.html", "testdata/expected_engineering_preflight.html"},
		{"testdata/research_article.html", "testdata/expected_research_preflight.html"},
	}

	for _, tc := range cases {
		t.Run(tc.source, func(t *testing.T) {
			root := parseHTMLFixture(t, tc.source)
			data := np.Collect(root)
			if len(data) == 0 {
				t.Fatal("no preflight data")
			}

			stream, err := np.Parse(data)
			if err != nil {
				t.Fatal(err)
			}
			if len(stream.Rows) == 0 {
				t.Fatal("no flight rows")
			}

			article := preflight.RenderArticleHTML(data)
			if article == nil {
				t.Fatal("renderer returned nil")
			}

			var content strings.Builder
			if err := html.Render(&content, article); err != nil {
				t.Fatal(err)
			}
			assertRenderedHTML(t, content.String(), tc.expected)
		})
	}
}
