package models

import (
	"github.com/imcrazytwkr/feedhub/models"
)

// Locale is a Cursor URL path segment
type Locale string

const (
	LocaleEn = Locale("")
	LocaleFr = Locale(models.LanguageFr)
	LocaleDe = Locale(models.LanguageDe)
	LocaleJa = Locale(models.LanguageJa)
	LocaleKo = Locale(models.LanguageKo)
	LocalePt = Locale(models.LanguagePt)
	LocaleEs = Locale(models.LanguageEs)
	LocaleRu = Locale(models.LanguageRu)
	LocaleId = Locale(models.LanguageId)
	LocaleTr = Locale(models.LanguageTr)
	LocaleZh = Locale("cn")
)

func (l Locale) String() string {
	return string(l)
}

const DefaultLocale = LocaleEn

var langToLocale = map[models.Language]Locale{
	models.LanguageEn: LocaleEn,
	models.LanguageFr: LocaleFr,
	models.LanguageDe: LocaleDe,
	models.LanguageJa: LocaleJa,
	models.LanguageKo: LocaleKo,
	models.LanguagePt: LocalePt,
	models.LanguageEs: LocaleEs,
	models.LanguageRu: LocaleRu,
	models.LanguageId: LocaleId,
	models.LanguageTr: LocaleTr,
	models.LanguageZh: LocaleZh,
}

func GetLocaleFor(lang models.Language) (locale Locale, ok bool) {
	locale, ok = langToLocale[lang]
	if !ok {
		locale = DefaultLocale
	}

	return
}
