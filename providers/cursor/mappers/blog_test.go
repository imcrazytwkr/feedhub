package mappers_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/providers/cursor/mappers"
	"github.com/imcrazytwkr/feedhub/utils/testutil"
	"golang.org/x/net/html"
)

func parseHTMLFixture(t *testing.T, path string) *html.Node {
	t.Helper()

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	root, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if root == nil {
		t.Fatal("parsed document is nil")
	}
	return root
}

func TestPluckBlogEntries(t *testing.T) {
	entries := mappers.PluckBlogEntries(parseHTMLFixture(t, "testdata/blog_list.html"))
	if len(entries) < 1 {
		t.Fatal("No entries were parsed")
	}

	var expected []*models.Entry
	err := testutil.ReadJson("testdata/expected_entries.json", &expected)
	if err != nil {
		t.Fatal(err)
	}

	if !testutil.SlicesEqual(entries, expected) {
		t.Log("slices aren't equal")
		t.FailNow()
	}
}

func TestPluckBlogDescription(t *testing.T) {
	root := parseHTMLFixture(t, "testdata/blog_list.html")
	got := mappers.PluckBlogDescription(root)
	want := "Latest updates and insights from the Cursor team. Learn about AI-powered coding, product updates, and development tips."
	if got != want {
		t.Fatalf("description = %q, want %q", got, want)
	}
}

func TestBlogSiteMeta(t *testing.T) {
	feed := mappers.BlogSiteMeta("", models.LanguageEn, "")
	if feed.Title != "Cursor Blog" {
		t.Errorf("title = %q", feed.Title)
	}
	if feed.Link != "https://cursor.com/blog" {
		t.Errorf("link = %q", feed.Link)
	}
	if feed.Language != "en" {
		t.Errorf("language = %q", feed.Language)
	}

	product := mappers.BlogSiteMeta("product", models.LanguageJa, "desc")
	if product.Title != "Cursor Blog (product)" {
		t.Errorf("product title = %q", product.Title)
	}
	if product.Link != "https://cursor.com/ja/blog/topic/product" {
		t.Errorf("product link = %q", product.Link)
	}
	if product.Language != "ja" {
		t.Errorf("product language = %q", product.Language)
	}
	if product.Description != "desc" {
		t.Errorf("product description = %q", product.Description)
	}

	zh := mappers.BlogSiteMeta("", models.LanguageZh, "")
	if zh.Link != "https://cursor.com/cn/blog" {
		t.Errorf("zh link = %q", zh.Link)
	}
	if zh.Language != "zh" {
		t.Errorf("zh language = %q", zh.Language)
	}
}

func TestPluckArticle(t *testing.T) {
	root := parseHTMLFixture(t, "testdata/blog_article.html")

	author := mappers.PluckArticleAuthor(root)
	if author != "Vicent Martí" {
		t.Fatalf("author = %q", author)
	}

	content, err := mappers.PluckArticleContent(root)
	if err != nil {
		t.Fatal(err)
	}

	actualDom, err := html.Parse(strings.NewReader(content))
	if err != nil {
		t.Fatal(err)
	}

	expectedContent, err := os.ReadFile("testdata/expected_content.html")
	if err != nil {
		t.Fatal(err)
	}

	expectedDom, err := html.Parse(bytes.NewReader(expectedContent))
	if err != nil {
		t.Fatal(err)
	}

	if !testutil.DomsEqual(expectedDom, actualDom) {
		t.Logf("invalid contents:\n%s\n", content)
		t.Logf("expected:\n%s\n", expectedContent)
		t.FailNow()
	}
}
