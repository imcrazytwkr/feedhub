package models

type Language string

const (
	LanguageUnknown = Language("")
	LanguageEn      = Language("en")
	LanguageJp      = Language("ja")
)

func (l Language) String() string {
	return string(l)
}

func ParseLanguage(value string) Language {
	switch value {
	case LanguageEn.String():
		return LanguageEn
	case LanguageJp.String():
		return LanguageJp
	default:
		return LanguageUnknown
	}
}
