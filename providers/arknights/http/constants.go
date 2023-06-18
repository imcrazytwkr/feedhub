package http

import "github.com/imcrazytwkr/feedhub/providers/arknights/models"

// @source: https://arknights.global/news
const hostPrefixGlobal = "https://arknights.global:10024"

var hostPrefixes = map[models.Language]string{
	models.LanguageEn: hostPrefixGlobal,

	// @source: https://www.arknights.jp/news
	models.LanguageJp: "https://www.arknights.jp:10014",
}

const newsPath = "/news"

// @NOTE: source limits count non-inclusively, meanine that
// for limit == 19, only 18 entries will be returned
const entryLimit = 19
