package mappers

import (
	"slices"
	"strings"

	"github.com/imcrazytwkr/feedhub/models"
	m "github.com/imcrazytwkr/feedhub/providers/anthropic/models"
)

const entryLimit = 50

func mapPostsToEntries(posts []*m.PublicationPost, path string) []*models.Entry {
	size := min(len(posts), entryLimit)
	if size == 0 {
		return nil
	}

	sortPosts(posts)

	entries := make([]*models.Entry, size)
	i := 0

	for _, post := range posts {
		slug := strings.TrimSpace(post.Slug.Value)
		title := strings.TrimSpace(post.Title)
		if len(slug) == 0 || len(title) == 0 {
			continue
		}

		entries[i] = &models.Entry{
			Title:       title,
			Link:        host + path + "/" + slug,
			Published:   post.PublishedAt,
			Description: strings.TrimSpace(post.Summary),
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

	//	No need for additional sort since posts are pre-sorted
	return entries
}

func sortPosts(posts []*m.PublicationPost) []*m.PublicationPost {
	if len(posts) > 1 {
		slices.SortFunc(posts, func(a, b *m.PublicationPost) int {
			return b.PublishedAt.Compare(a.PublishedAt)
		})
	}

	return posts
}
