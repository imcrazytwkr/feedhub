package mappers

import "github.com/imcrazytwkr/feedhub/providers/arknights/models"

const hostPrefixGlobal = "https://arknights.global/news/"

var hostPrefixes = map[models.Language]string{
	models.LanguageEn: hostPrefixGlobal,
	models.LanguageJp: "https://www.arknights.jp/news/",
}

var feedTitles = map[models.Language]string{
	models.LanguageEn: "Arknights",
	models.LanguageJp: "アークナイツ",
}

var feedDescriptions = map[models.Language]string{
	models.LanguageEn: "Arknights news",
	models.LanguageJp: "アークナイツ ニュース",
}
