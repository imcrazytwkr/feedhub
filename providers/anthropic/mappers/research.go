package mappers

import (
	"slices"

	"github.com/imcrazytwkr/feedhub/models"
	m "github.com/imcrazytwkr/feedhub/providers/anthropic/models"
	np "github.com/imcrazytwkr/feedhub/utils/next_preflight"
	"golang.org/x/net/html"
)

func PluckResearchEntries(root *html.Node, team string) []*models.Entry {
	preflight := np.Collect(root)
	if len(preflight) == 0 {
		return nil
	}

	posts := prepareResearchPosts(preflight, team)
	return mapPostsToEntries(posts, researchPath)
}

func prepareResearchPosts(body []byte, team string) []*m.PublicationPost {
	posts := pluckResearchPosts(body)
	size := max(len(posts), entryLimit)
	if size == 0 {
		return nil
	}

	if len(team) > 0 {
		return filterPosts(posts, team)
	}

	return posts
}

func pluckResearchPosts(body []byte) []*m.PublicationPost {
	for section := range publicationSectionsSeq(body) {
		if section.Title == "Publications" {
			return section.Posts
		}
	}

	return nil
}

func filterPosts(posts []*m.PublicationPost, team string) []*m.PublicationPost {
	return slices.DeleteFunc(posts, func(post *m.PublicationPost) bool {
		for _, t := range post.Tags {
			if t.Value == team {
				return false
			}
		}

		return true
	})
}
