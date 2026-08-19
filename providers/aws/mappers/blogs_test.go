package mappers_test

import (
	"testing"

	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/providers/aws/mappers"
	m "github.com/imcrazytwkr/feedhub/providers/aws/models"
	"github.com/imcrazytwkr/feedhub/utils/testutil"
)

func TestBlogsEntriesExtraction(t *testing.T) {
	var contents m.SearchResponse
	err := testutil.ReadJson("testdata/blogs_all_en.json", &contents)
	if err != nil {
		t.Fatal(err)
	}

	entries := mappers.PluckEntries(&contents)
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

func TestSiteMetaUnfilteredEnglish(t *testing.T) {
	actual := mappers.GenerateSiteMeta(models.LanguageEn, "")

	var expected models.Feed
	err := testutil.ReadJson("testdata/expected_meta.json", &expected)
	if err != nil {
		t.Fatal(err)
	}

	assertSiteMeta(t, actual, &expected)
}

func TestGenerateSiteMeta(t *testing.T) {
	cases := []struct {
		language models.Language
		category string
		title    string
		link     string
	}{
		{models.LanguageEn, "", "AWS Blog", "https://aws.amazon.com/blogs/"},
		{models.LanguageEn, "database", "AWS Blog (database)", "https://aws.amazon.com/blogs/database/"},
		{models.LanguageJa, "", "AWS Blog", "https://aws.amazon.com/ja/blogs/"},
		{models.LanguageJa, "database", "AWS Blog (database)", "https://aws.amazon.com/ja/blogs/database/"},
		{models.LanguageZh, "compute", "AWS Blog (compute)", "https://aws.amazon.com/zh/blogs/compute/"},
		{models.LanguageFr, "", "AWS Blog", "https://aws.amazon.com/fr/blogs/"},
	}

	for _, tc := range cases {
		name := tc.language.String()
		if len(tc.category) > 0 {
			name += "_" + tc.category
		}

		t.Run(name, func(t *testing.T) {
			actual := mappers.GenerateSiteMeta(tc.language, tc.category)
			expected := &models.Feed{
				Title:    tc.title,
				Language: tc.language.String(),
				Link:     tc.link,
			}
			assertSiteMeta(t, actual, expected)
		})
	}
}

func assertSiteMeta(t *testing.T, actual *models.Feed, expected *models.Feed) {
	t.Helper()

	if actual.Title != expected.Title {
		t.Errorf("Titles mismatch, expected %q, got %q", expected.Title, actual.Title)
	}

	if len(expected.Description) > 0 && actual.Description != expected.Description {
		t.Errorf("Descriptions mismatch, expected %q, got %q", expected.Description, actual.Description)
	}

	if actual.Language != expected.Language {
		t.Errorf("Language mismatch, expected %q, got %q", expected.Language, actual.Language)
	}

	if actual.Link != expected.Link {
		t.Errorf("Link mismatch, expected %q, got %q", expected.Link, actual.Link)
	}
}
