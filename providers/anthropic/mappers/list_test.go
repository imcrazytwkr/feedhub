package mappers_test

import (
	"testing"

	"github.com/imcrazytwkr/feedhub/providers/anthropic/mappers"
)

func TestNewsPreflightEntriesExtraction(t *testing.T) {
	assertEntries(
		t,
		mappers.PluckNewsEntries(parseHTMLFixture(t, "testdata/news_list.html")),
		"testdata/expected_news_preflight_entries.json",
	)
}

func TestEngineeringPreflightEntriesExtraction(t *testing.T) {
	assertEntries(
		t,
		mappers.PluckEngineeringEntries(parseHTMLFixture(t, "testdata/engineering_list.html")),
		"testdata/expected_engineering_preflight_entries.json",
	)
}

func TestResearchEntriesExtraction(t *testing.T) {
	assertEntries(
		t,
		mappers.PluckResearchEntries(parseHTMLFixture(t, "testdata/research_list.html"), ""),
		"testdata/expected_research_entries.json",
	)
}

func TestResearchTeamFilter(t *testing.T) {
	root := parseHTMLFixture(t, "testdata/research_list.html")
	assertEntries(t, mappers.PluckResearchEntries(root, "alignment"), "testdata/expected_research_alignment_entries.json")

	none := mappers.PluckResearchEntries(root, "missing-team")
	if none != nil {
		t.Fatalf("missing team should yield nil, got %d entries", len(none))
	}
}
