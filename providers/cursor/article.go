package cursor

import (
	"bytes"
	"context"

	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/imcrazytwkr/feedhub/providers/cursor/mappers"
	"github.com/rs/zerolog"
	"golang.org/x/net/html"
)

func (p *cursorProvider) getArticleData(ctx context.Context, articleURL string) (author, content string, err error) {
	log := zerolog.Ctx(ctx).With().Str("url", articleURL).Logger()

	body, err := p.client.FetchArticle(ctx, articleURL)
	if err != nil {
		log.Debug().Err(err).Msg("failed to fetch article")
		return "", "", err
	}

	article, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		log.Debug().Err(err).Msg("failed to parse article")
		return "", "", constants.ErrorMalformedBody
	}

	content, err = mappers.PluckArticleContent(article)
	if err != nil {
		log.Debug().Err(err).Msg("error rendering article")
		return "", "", constants.ErrorMalformedBody
	}

	return mappers.PluckArticleAuthor(article), content, nil
}
