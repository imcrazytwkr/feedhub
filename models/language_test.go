package models

import "testing"

func TestParseLanguage(t *testing.T) {
	cases := []struct {
		in   string
		want Language
	}{
		{"", LanguageUnknown},
		{"en", LanguageEn},
		{"fr", LanguageFr},
		{"de", LanguageDe},
		{"ja", LanguageJa},
		{"ko", LanguageKo},
		{"pt", LanguagePt},
		{"es", LanguageEs},
		{"ru", LanguageRu},
		{"id", LanguageId},
		{"tr", LanguageTr},
		{"zh", LanguageZh},
		{"EN", LanguageUnknown},
		{"en_US", LanguageUnknown},
		{"jp", LanguageUnknown},
		{"database", LanguageUnknown},
	}

	for _, tc := range cases {
		got := ParseLanguage(tc.in)
		if got != tc.want {
			t.Errorf("ParseLanguage(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}
