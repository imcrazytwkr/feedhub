package mappers_test

import (
	"os"
	"testing"

	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/providers/arknights/mappers"
	m "github.com/imcrazytwkr/feedhub/providers/arknights/models"
	"github.com/imcrazytwkr/feedhub/utils/testutil"
	"github.com/valyala/fastjson"
)

func TestNewsEntriesExtraction(t *testing.T) {
	feedFile, err := os.ReadFile("test_data/news_feed.json")
	if err != nil {
		t.Fatal(err)
	}

	parser := &fastjson.Parser{}

	contents, err := parser.ParseBytes(feedFile)
	if err != nil {
		t.Fatal(err)
	}

	if contents == nil {
		t.Fatal("JSON data is NIL!")
	}

	entries := mappers.PluckEntries(contents, m.LanguageEn)
	if len(entries) < 1 {
		t.Fatal("No entries were parsed")
	}

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
