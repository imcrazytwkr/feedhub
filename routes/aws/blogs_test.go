package aws

import (
	"testing"

	"github.com/imcrazytwkr/feedhub/models"
)

func TestResolveBlogsPath(t *testing.T) {
	cases := []struct {
		first, second string
		wantCategory  string
		wantLang      models.Language
	}{
		{"", "", "", models.LanguageEn},
		{"ja", "", "", models.LanguageJa},
		{"zh", "", "", models.LanguageZh},
		{"database", "", "database", models.LanguageEn},
		{"database", "ja", "database", models.LanguageJa},
		{"id", "", "", models.LanguageId},
		{"id", "en", "id", models.LanguageEn},
	}

	for _, tc := range cases {
		category, lang := resolveBlogsPath(tc.first, tc.second)
		if category != tc.wantCategory || lang != tc.wantLang {
			t.Errorf("resolveBlogsPath(%q, %q) = (%q, %q), want (%q, %q)",
				tc.first, tc.second, category, lang, tc.wantCategory, tc.wantLang)
		}
	}
}
