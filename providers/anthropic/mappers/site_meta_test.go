package mappers_test

import (
	"testing"

	"github.com/imcrazytwkr/feedhub/providers/anthropic/mappers"
)

func TestSiteMeta(t *testing.T) {
	news := mappers.NewsSiteMeta()
	if news.Title != "Anthropic News" {
		t.Errorf("news title = %q", news.Title)
	}
	if news.Link != "https://www.anthropic.com/news" {
		t.Errorf("news link = %q", news.Link)
	}
	if news.Language != "en" {
		t.Errorf("news language = %q, want en", news.Language)
	}

	eng := mappers.EngineeringSiteMeta()
	if eng.Title != "Anthropic Engineering" {
		t.Errorf("engineering title = %q", eng.Title)
	}
	if eng.Link != "https://www.anthropic.com/engineering" {
		t.Errorf("engineering link = %q", eng.Link)
	}

	research := mappers.ResearchSiteMeta("")
	if research.Title != "Anthropic Research" {
		t.Errorf("research title = %q", research.Title)
	}
	if research.Link != "https://www.anthropic.com/research" {
		t.Errorf("research link = %q", research.Link)
	}

	alignment := mappers.ResearchSiteMeta("alignment")
	if alignment.Title != "Anthropic Research (alignment)" {
		t.Errorf("alignment title = %q", alignment.Title)
	}
	if alignment.Link != "https://www.anthropic.com/research/team/alignment" {
		t.Errorf("alignment link = %q", alignment.Link)
	}
}
