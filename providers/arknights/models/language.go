package models

type Language string

const (
	LanguageEn = Language("en")
	LanguageJp = Language("ja")
)

func (l Language) String() string {
	return string(l)
}

func ParseFormat(value string) Language {
	if value == string(LanguageJp) {
		return LanguageJp
	}

	return LanguageEn
}
