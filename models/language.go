package models

type Language string

const (
	LanguageUnknown = Language("")
	LanguageEn      = Language("en")
	LanguageFr      = Language("fr")
	LanguageDe      = Language("de")
	LanguageJa      = Language("ja")
	LanguageKo      = Language("ko")
	LanguagePt      = Language("pt")
	LanguageEs      = Language("es")
	LanguageRu      = Language("ru")
	LanguageId      = Language("id")
	LanguageTr      = Language("tr")
	LanguageZh      = Language("zh")
)

func (l Language) String() string {
	return string(l)
}

func ParseLanguage(value string) Language {
	lang := Language(value)
	switch lang {
	case LanguageEn,
		LanguageFr,
		LanguageDe,
		LanguageJa,
		LanguageKo,
		LanguagePt,
		LanguageEs,
		LanguageRu,
		LanguageId,
		LanguageTr,
		LanguageZh:
		return lang
	default:
		return LanguageUnknown
	}
}
