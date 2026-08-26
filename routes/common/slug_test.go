package common

import (
	"testing"

	"github.com/imcrazytwkr/feedhub/models"
)

func TestParseSlugWithLocale(t *testing.T) {
	cases := []struct {
		first, second string
		wantSlug      string
		wantLang      models.Language
	}{
		{"", "", "", models.LanguageEn},
		{"ja", "", "", models.LanguageJa},
		{"zh", "", "", models.LanguageZh},
		{"product", "", "product", models.LanguageEn},
		{"database", "", "database", models.LanguageEn},
		{"product", "ja", "product", models.LanguageJa},
		{"database", "ja", "database", models.LanguageJa},
		{"id", "", "", models.LanguageId},
		{"id", "en", "id", models.LanguageEn},
		{"product", "nope", "product", models.LanguageEn},
		{"en", "ja", "en", models.LanguageJa},
	}

	for _, tc := range cases {
		slug, lang := ParseSlugWithLocale(tc.first, tc.second)
		if slug != tc.wantSlug || lang != tc.wantLang {
			t.Errorf("ParseSlugWithLocale(%q, %q) = (%q, %q), want (%q, %q)",
				tc.first, tc.second, slug, lang, tc.wantSlug, tc.wantLang)
		}
	}
}
