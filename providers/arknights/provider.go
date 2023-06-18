package arknights

import (
	"context"
	"net/http"

	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/providers"
	h "github.com/imcrazytwkr/feedhub/providers/arknights/http"
	"github.com/imcrazytwkr/feedhub/providers/arknights/mappers"
	m "github.com/imcrazytwkr/feedhub/providers/arknights/models"
	"github.com/rs/zerolog"
	"github.com/valyala/fastjson"
)

type arknightsProvider struct {
	parserPool *fastjson.ParserPool
	client     h.ArknightsClient
}

func NewArknightsProvider(parserPool *fastjson.ParserPool, httpClient *http.Client) providers.ArknightsProvider {
	return &arknightsProvider{
		parserPool: parserPool,
		client:     *h.NewArknightsClient(httpClient),
	}
}

func (p *arknightsProvider) GetNews(ctx context.Context, lang string) (*models.Feed, error) {
	log := zerolog.Ctx(ctx)

	language := m.ParseLanguage(lang)
	if language == m.LanguageUnknown {
		log.Debug().Msgf("unknown language: %q", lang)
		return nil, nil
	}

	// Entries from RSS
	body, err := p.client.GetNews(ctx, language)
	if err != nil {
		log.Debug().Err(err).Str("lang", language.String()).Msg("failed to fetch news")
		return nil, err
	}

	log.Trace().Str("lang", language.String()).Msg("fetched news RSS")

	parser := p.parserPool.Get()
	defer p.parserPool.Put(parser)

	payload, err := parser.ParseBytes(body)
	if err != nil {
		log.Debug().Err(err).Str("lang", language.String()).Msg("failed to parse news")
		return nil, constants.ErrorMalformedBody
	}

	log.Trace().Str("lang", language.String()).Msg("successfully parsed news")

	entries := mappers.PluckEntries(payload, language)
	if len(entries) == 0 {
		log.Trace().Msg("no news found, exiting early")
		return nil, nil
	}

	feed := mappers.GenerateSiteMeta(language)
	feed.Published = entries[0].Published
	feed.Entries = entries
	return feed, nil
}
