package mappers

import (
	"bytes"
	"encoding/json"

	"github.com/imcrazytwkr/feedhub/models"
	m "github.com/imcrazytwkr/feedhub/providers/anthropic/models"
	np "github.com/imcrazytwkr/feedhub/utils/next_preflight"
	"golang.org/x/net/html"
)

func PluckEngineeringEntries(root *html.Node) []*models.Entry {
	preflight := np.Collect(root)
	if len(preflight) == 0 {
		return nil
	}

	articles := pluckEngineeringArticles(preflight)
	return mapArticlesToEntries(articles)
}

func pluckEngineeringArticles(body []byte) []*m.Article {
	for line := range bytes.Lines(body) {
		if !bytes.Contains(line, []byte("articleList")) {
			continue
		}

		s := bytes.IndexByte(line, '[')
		if s < 0 {
			continue
		}

		var row []json.RawMessage
		if json.Unmarshal(line[s:], &row) != nil {
			continue
		}

		for _, item := range row {
			if len(item) == 0 || item[0] != '{' || item[len(item)-1] != '}' {
				continue
			}

			var container m.Container[m.ArticleSection]
			if json.Unmarshal(item, &container) != nil {
				continue
			}

			if container.Page == nil || len(container.Page.Sections) == 0 {
				continue
			}

			for i := range container.Page.Sections {
				if len(container.Page.Sections[i].Articles) > 0 {
					return container.Page.Sections[i].Articles
				}
			}
		}
	}

	return nil
}
