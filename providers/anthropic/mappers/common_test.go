package mappers_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/imcrazytwkr/feedhub/models"
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

func assertEntries(t *testing.T, entries []*models.Entry, expectedPath string) {
	t.Helper()

	if len(entries) < 1 {
		t.Fatal("No entries were parsed")
	}

	var expected []*models.Entry
	err := testutil.ReadJson(expectedPath, &expected)
	if err != nil {
		t.Fatal(err)
	}

	if !testutil.SlicesEqual(entries, expected) {
		t.Log("slices aren't equal")
		t.FailNow()
	}
}

func assertRenderedHTML(t *testing.T, content, expectedPath string) {
	t.Helper()

	actualDom, err := html.Parse(strings.NewReader(content))
	if err != nil {
		t.Fatal(err)
	}

	expectedContent, err := os.ReadFile(expectedPath)
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
