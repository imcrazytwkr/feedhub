package mappers

import "github.com/imcrazytwkr/feedhub/models"

const hostPrefixGlobal = "https://arknights.global/news/"

var hostPrefixes = map[models.Language]string{
	models.LanguageEn: hostPrefixGlobal,
	models.LanguageJa: "https://www.arknights.jp/news/",
}

var feedTitles = map[models.Language]string{
	models.LanguageEn: "Arknights",
	models.LanguageJa: "アークナイツ",
}

var feedDescriptions = map[models.Language]string{
	models.LanguageEn: "Arknights news",
	models.LanguageJa: "アークナイツ ニュース",
}
