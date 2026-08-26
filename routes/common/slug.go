package common

import "github.com/imcrazytwkr/feedhub/models"

/**
 * Implements "optional lang tail segment" parsing
 *
 * These cases are excessively documented because nested optional params
 * are terrible from code standpoint but good for UX.
 */
func ParseSlugWithLocale(category string, lang string) (string, models.Language) {
	// `//:lang` is normalized to `/:lang` by GIN so it is impossible for "lang"
	// to be set when "topic" is empty.
	if len(category) == 0 {
		// No parameters (all topics in English)
		return "", models.LanguageEn
	}

	if len(lang) == 0 {
		language := models.ParseLanguage(category)

		// `/provider/:lang` (all posts in language)
		if language != models.LanguageUnknown {
			return "", language
		}

		// `/provider/:slug` (posts for category in English)
		return category, models.LanguageEn
	}

	// `/provider/:slug/:language` (posts for topic in language)
	language := models.ParseLanguage(lang)
	if language == models.LanguageUnknown {
		return category, models.LanguageEn
	}

	return category, language
}
