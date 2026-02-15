package mappers_test

import (
	"os"
	"testing"
	"time"

	"github.com/antchfx/xmlquery"
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/providers/tomrigby/mappers"
	"github.com/imcrazytwkr/feedhub/utils/testutil"
)

func TestBlogEntriesExtraction(t *testing.T) {
	feedFile, err := os.Open("test_data/blog_feed.xml")
	if err != nil {
		t.Fatal(err)
	}

	contents, err := xmlquery.Parse(feedFile)
	if err != nil {
		t.Fatal(err)
	}

	err = feedFile.Close()
	if err != nil {
		t.Fatal(err)
	}

	if contents == nil {
		t.Fatal("rss feed is NIL!")
	}

	entries, err := mappers.PluckEntries(contents)
	if err != nil {
		t.Fatal(err)
	}

	if len(entries) == 0 {
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

func TestSiteMetaExtraction(t *testing.T) {
	feedFile, err := os.Open("test_data/blog_feed.xml")
	if err != nil {
		t.Fatal(err)
	}

	contents, err := xmlquery.Parse(feedFile)
	if err != nil {
		t.Fatal(err)
	}

	err = feedFile.Close()
	if err != nil {
		t.Fatal(err)
	}

	if contents == nil {
		t.Fatal("rss feed is NIL!")
	}

	actualMeta := mappers.PickSiteMeta(contents)

	var expectedMeta models.Feed
	testutil.ReadJson("test_data/expected_meta.json", &expectedMeta)

	if actualMeta.Title != expectedMeta.Title {
		t.Logf("Titles mismatch, expected %q, got %q", expectedMeta.Title, actualMeta.Title)
		t.Fail()
	}

	if actualMeta.Description != expectedMeta.Description {
		t.Logf("Descriptions mismatch, expected %q, got %q", expectedMeta.Description, actualMeta.Description)
		t.Fail()
	}

	if actualMeta.Language != expectedMeta.Language {
		t.Logf("Language mismatch, expected %q, got %q", expectedMeta.Language, actualMeta.Language)
		t.Fail()
	}

	if !actualMeta.Published.Equal(expectedMeta.Published) {
		t.Logf("Published mismatch, expected %q, got %q", expectedMeta.Published.Format(time.RFC3339), actualMeta.Published.Format(time.RFC3339))
		t.Fail()
	}
}
