package xmlutil_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/antchfx/xmlquery"
	"github.com/imcrazytwkr/feedhub/utils/testutil"
	"github.com/imcrazytwkr/feedhub/utils/xmlutil"
	"golang.org/x/net/html"
)

func TestBlogEntriesExtraction(t *testing.T) {
	feedFile, err := os.Open("testdata/malformed_content.xml")
	if err != nil {
		t.Fatal(err)
	}

	actualRoot, err := xmlquery.Parse(feedFile)
	if err != nil {
		t.Fatal(err)
	}

	err = feedFile.Close()
	if err != nil {
		t.Fatal(err)
	}

	if actualRoot == nil {
		t.Fatal("Source XML file is empty!")
	}

	entryContents := xmlquery.FindOne(actualRoot, "//content")
	if entryContents == nil {
		t.Fatal("Malformed source data XML (no <content> tag)")
	}

	var sb strings.Builder

	for child := entryContents.FirstChild; child != nil; child = child.NextSibling {
		err = xmlutil.RenderNode(&sb, child)
		if err != nil {
			t.Fatal(err)
		}
	}

	actualContent := strings.TrimSpace(sb.String())

	actualDom, err := html.Parse(strings.NewReader(actualContent))
	if err != nil {
		t.Fatal(err)
	}

	if actualDom == nil {
		t.Fatalf("Could not parse actual DOM from actual content:\n%s\n", actualContent)
	}

	expectedContent, err := os.ReadFile("testdata/clean_content.html")
	if err != nil {
		t.Fatal(err)
	}

	expectedDom, err := html.Parse(bytes.NewReader(expectedContent))
	if err != nil {
		t.Fatal(err)
	}

	if expectedDom == nil {
		t.Fatalf("Could not parse expected DOM from expected content:\n%s\n", expectedContent)
	}

	if !testutil.DomsEqual(expectedDom, actualDom) {
		t.Logf("Invalud contents:\n%s\n", actualContent)
		t.Logf("Expected:\n%s\n", expectedContent)
		t.FailNow()
	}
}
