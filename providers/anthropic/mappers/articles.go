package mappers

import (
	"slices"
	"strings"

	"github.com/imcrazytwkr/feedhub/models"
	m "github.com/imcrazytwkr/feedhub/providers/anthropic/models"
)

func mapArticlesToEntries(articles []*m.Article) []*models.Entry {
	size := min(len(articles), entryLimit)
	if size == 0 {
		return nil
	}

	sortArticles(articles)

	entries := make([]*models.Entry, size)
	i := 0

	for _, article := range articles {
		slug := strings.TrimSpace(article.Slug.Value)
		title := strings.TrimSpace(article.Title)
		if len(slug) == 0 || len(title) == 0 {
			return nil
		}

		entries[i] = &models.Entry{
			Title:       title,
			Link:        host + engineeringPath + "/" + slug,
			Published:   article.PublishedAt.Time(),
			Description: strings.TrimSpace(article.Summary),
		}

		i++

		if i >= size {
			break
		}
	}

	if i == 0 {
		return nil
	}

	if i < size {
		return entries[0:i]
	}

	//	No need for additional sort since articles are pre-sorted
	return entries
}

func sortArticles(articles []*m.Article) []*m.Article {
	if len(articles) > 1 {
		slices.SortFunc(articles, func(a, b *m.Article) int {
			return b.PublishedAt.Time().Compare(a.PublishedAt.Time())
		})
	}

	return articles
}
