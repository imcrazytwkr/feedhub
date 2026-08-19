package models

import (
	"github.com/imcrazytwkr/feedhub/models"
)

type Locale string

const (
	LocaleEn = Locale("en_US")
	LocaleFr = Locale("fr_FR")
	LocaleDe = Locale("de_DE")
	LocaleJp = Locale("ja_JP")
	LocaleKo = Locale("ko_KR")
	LocalePt = Locale("pt_BR")
	LocaleEs = Locale("es_ES")
	LocaleRu = Locale("ru_RU")
	LocaleId = Locale("id_ID")
	LocaleTr = Locale("tr_TR")
	LocaleZh = Locale("zh_CN")
)

func (l Locale) String() string {
	return string(l)
}

const DefaultLocale = LocaleEn

var langToLocale = map[models.Language]Locale{
	models.LanguageEn: LocaleEn,
	models.LanguageFr: LocaleFr,
	models.LanguageDe: LocaleDe,
	models.LanguageJa: LocaleJp,
	models.LanguageKo: LocaleKo,
	models.LanguagePt: LocalePt,
	models.LanguageEs: LocaleEs,
	models.LanguageRu: LocaleRu,
	models.LanguageId: LocaleId,
	models.LanguageTr: LocaleTr,
	models.LanguageZh: LocaleZh,
}

func GetLocaleFor(lang models.Language) Locale {
	locale, ok := langToLocale[lang]
	if !ok {
		return DefaultLocale
	}

	return locale
}
