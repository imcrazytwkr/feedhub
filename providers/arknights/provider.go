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

func (p *arknightsProvider) GetNews(ctx context.Context, lang models.Language) (*models.Feed, error) {
	switch lang {
	case models.LanguageEn, models.LanguageJa:
		// All good, nothing to do here
		break
	default:
		zerolog.Ctx(ctx).Debug().Msgf("unsupported language: %q", lang)
		return nil, nil
	}

	log := zerolog.Ctx(ctx).With().Str("lang", lang.String()).Logger()

	// Entries from API
	body, err := p.client.GetNews(ctx, lang)
	if err != nil {
		log.Debug().Err(err).Msg("failed to fetch news")
		return nil, err
	}

	log.Trace().Msg("fetched news")

	var payload m.NewsResponse
	err = json.Unmarshal(body, &payload)
	if err != nil {
		log.Debug().Err(err).Msg("failed to parse news")
		return nil, constants.ErrorMalformedBody
	}

	log.Trace().Msg("successfully parsed news")

	entries := mappers.PluckEntries(&payload, lang)
	if len(entries) == 0 {
		log.Trace().Msg("no news found, exiting early")
		return nil, nil
	}

	feed := mappers.GenerateSiteMeta(lang)
	feed.Published = entries[0].Published
	feed.Entries = entries
	return feed, nil
}
