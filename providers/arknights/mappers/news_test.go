package mappers_test

import (
	"testing"

	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/providers/arknights/mappers"
	m "github.com/imcrazytwkr/feedhub/providers/arknights/models"
	"github.com/imcrazytwkr/feedhub/utils/testutil"
)

func TestNewsEntriesExtraction(t *testing.T) {
	var contents m.NewsResponse
	err := testutil.ReadJson("testdata/news_feed.json", &contents)
	if err != nil {
		t.Fatal(err)
	}

	entries := mappers.PluckEntries(&contents, models.LanguageEn)
	if len(entries) < 1 {
		t.Fatal("No entries were parsed")
	}

	var expectedEntries []*models.Entry
	err = testutil.ReadJson("testdata/expected_entries.json", &expectedEntries)
	if err != nil {
		t.Fatal(err)
	}

	if !testutil.SlicesEqual(entries, expectedEntries) {
		t.Log("slices aren't equal")
		t.FailNow()
	}
}
