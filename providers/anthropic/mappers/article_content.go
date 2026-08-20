package mappers

import (
	"strings"

	p "github.com/imcrazytwkr/feedhub/providers/anthropic/mappers/preflight"
	np "github.com/imcrazytwkr/feedhub/utils/next_preflight"
	"golang.org/x/net/html"
)

func PluckArticleContent(root *html.Node) (string, error) {
	preflight := np.Collect(root)
	if len(preflight) == 0 {
		return "", nil
	}

	node := p.RenderArticleHTML(preflight)
	if node == nil {
		return "", nil
	}

	var sb strings.Builder
	err := html.Render(&sb, node)
	if err != nil {
		return "", err
	}

	return sb.String(), nil
}
