package mappers

import (
	"github.com/imcrazytwkr/feedhub/models"
	m "github.com/imcrazytwkr/feedhub/providers/anthropic/models"
	np "github.com/imcrazytwkr/feedhub/utils/next_preflight"
	"golang.org/x/net/html"
)

func PluckNewsEntries(root *html.Node) []*models.Entry {
	preflight := np.Collect(root)
	if len(preflight) == 0 {
		return nil
	}

	posts := pluckNewsPosts(preflight)
	return mapPostsToEntries(posts, newsPath)
}

func pluckNewsPosts(body []byte) []*m.PublicationPost {
	for section := range publicationSectionsSeq(body) {
		if section.Title == "News" {
			return section.Posts
		}
	}

	return nil
}
