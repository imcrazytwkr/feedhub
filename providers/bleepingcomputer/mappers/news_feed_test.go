package mappers_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/antchfx/xmlquery"
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/providers/bleepingcomputer/mappers"
	"github.com/imcrazytwkr/feedhub/utils/testutil"
)

func TestNewsItemsExtraction(t *testing.T) {
	feedFile, err := os.ReadFile("test_data/news_feed.xml")
	if err != nil {
		t.Fatal(err)
	}

	contents, err := xmlquery.Parse(bytes.NewReader(feedFile))
	if err != nil {
		t.Fatal(err)
	}

	if contents == nil {
		t.Fatal("rss feed is NIL!")
	}

	entries := mappers.PluckItems(contents)
	if len(entries) < 1 {
		t.Fatal("No entries were parsed")
	}

	fmt.Printf("%v", entries)

	var expectedEntries []*models.Entry
	err = testutil.ReadJson("test_data/expected_entries.json", &expectedEntries)
	if err != nil {
		t.Fatal(err)
	}

	if !testutil.SlicesEqual(entries, expectedEntries) {
		t.Log("slices aren't equal")
		t.FailNow()
	}
}
