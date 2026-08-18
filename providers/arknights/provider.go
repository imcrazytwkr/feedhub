package arknights

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/providers"
	h "github.com/imcrazytwkr/feedhub/providers/arknights/http"
	"github.com/imcrazytwkr/feedhub/providers/arknights/mappers"
	m "github.com/imcrazytwkr/feedhub/providers/arknights/models"
	"github.com/rs/zerolog"
)

type arknightsProvider struct {
	client *h.ArknightsClient
}

func NewArknightsProvider(httpClient *http.Client) providers.ArknightsProvider {
	return &arknightsProvider{h.NewArknightsClient(httpClient)}
}

func (p *arknightsProvider) GetNews(ctx context.Context, lang string) (*models.Feed, error) {
	log := zerolog.Ctx(ctx)

	language := m.ParseLanguage(lang)
	if language == m.LanguageUnknown {
		log.Debug().Msgf("unknown language: %q", lang)
		return nil, nil
	}

	// Entries from API
	body, err := p.client.GetNews(ctx, language)
	if err != nil {
		log.Debug().Str("lang", lang).Err(err).Msg("failed to fetch news")
		return nil, err
	}

	log.Trace().Str("lang", lang).Msg("fetched news")

	var payload m.NewsResponse
	err = json.Unmarshal(body, &payload)
	if err != nil {
		log.Debug().Err(err).Str("lang", lang).Msg("failed to parse news")
		return nil, constants.ErrorMalformedBody
	}

	log.Trace().Str("lang", lang).Msg("successfully parsed news")

	entries := mappers.PluckEntries(&payload, language)
	if len(entries) == 0 {
		log.Trace().Str("lang", lang).Msg("no news found, exiting early")
		return nil, nil
	}

	feed := mappers.GenerateSiteMeta(language)
	feed.Published = entries[0].Published
	feed.Entries = entries
	return feed, nil
}
