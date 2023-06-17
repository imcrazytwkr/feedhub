package mappers_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/antchfx/htmlquery"
	"github.com/imcrazytwkr/feedhub/providers/bleepingcomputer/mappers"
	"github.com/imcrazytwkr/feedhub/utils/testutil"
	"golang.org/x/net/html"
)

const expectedDescription = "The MOVEit Transfer extortion attacks continue to dominate the news cycle, with the Clop ransomware operation now extorting organizations breached in the attacks."

func TestNewsArticleDescriptionExtraction(t *testing.T) {
	articleFile, err := os.ReadFile("test_data/news_article.html")
	if err != nil {
		t.Fatal(err)
	}

	contents, err := htmlquery.Parse(bytes.NewReader(articleFile))
	if err != nil {
		t.Fatal(err)
	}

	if contents == nil {
		t.Fatal("rss feed is NIL!")
	}

	description := mappers.PickArticleDescription(contents)
	if description != expectedDescription {
		t.Logf("Invalid description: %q\n", description)
		t.Logf("Expected: %q\n", expectedDescription)
		t.FailNow()
	}
}

func TestNewsArticleContentExtraction(t *testing.T) {
	articleFile, err := os.ReadFile("test_data/news_article.html")
	if err != nil {
		t.Fatal(err)
	}

	contents, err := htmlquery.Parse(bytes.NewReader(articleFile))
	if err != nil {
		t.Fatal(err)
	}

	if contents == nil {
		t.Fatal("rss feed is NIL!")
	}

	articleContent, err := mappers.PickArticleContents(contents)
	if err != nil {
		t.Fatal(err)
	}

	actualleDom, err := html.Parse(strings.NewReader(articleContent))
	if err != nil {
		t.Fatal(err)
	}

	expectedContent, err := os.ReadFile("test_data/expected_content.html")
	if err != nil {
		t.Fatal(err)
	}

	expectedDom, err := html.Parse(bytes.NewReader(expectedContent))
	if err != nil {
		t.Fatal(err)
	}

	if !testutil.DomsEqual(expectedDom, actualleDom) {
		t.Logf("Invalud contents:\n%s\n", articleContent)
		t.Logf("Expected:\n%s\n", expectedContent)
		t.FailNow()
	}
}
