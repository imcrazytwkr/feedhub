package anthropic

import (
	"bytes"
	"context"

	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/imcrazytwkr/feedhub/providers/anthropic/mappers"
	"github.com/rs/zerolog"
	"golang.org/x/net/html"
)

func (p *anthropicProvider) getArticleContent(ctx context.Context, url string) (string, error) {
	log := zerolog.Ctx(ctx).With().Str("url", url).Logger()

	body, err := p.client.FetchArticle(ctx, url)
	if err != nil {
		log.Debug().Err(err).Msg("failed to fetch article")
		return "", err
	}

	article, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		log.Debug().Err(err).Msg("failed to parse article")
		return "", constants.ErrorMalformedBody
	}

	content, err := mappers.PluckArticleContent(article)
	if err != nil {
		log.Debug().Err(err).Msg("error rendering article")
		return "", constants.ErrorMalformedBody
	}

	return content, nil
}
