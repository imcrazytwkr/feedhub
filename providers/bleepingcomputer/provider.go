package bleepingcomputer

import (
	"bytes"
	"context"
	"net/http"

	"github.com/antchfx/htmlquery"
	"github.com/antchfx/xmlquery"
	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/providers"
	h "github.com/imcrazytwkr/feedhub/providers/bleepingcomputer/http"
	m "github.com/imcrazytwkr/feedhub/providers/bleepingcomputer/mappers"
	"github.com/rs/zerolog"
)

type bleepingComputerProvider struct {
	client *h.BleepingComputerClient
}

func NewBleepingComputerProvider(httpClient *http.Client) providers.BleepingComputerProvider {
	return &bleepingComputerProvider{
		client: h.NewPixivClient(httpClient),
	}
}

func (p *bleepingComputerProvider) GetNews(ctx context.Context) (*models.Feed, error) {
	log := zerolog.Ctx(ctx)

	// Entries from RSS
	body, err := p.client.FetchNewsFeed(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("failed to fetch news RSS")
		return nil, err
	}

	log.Trace().Msg("fetched news RSS")

	root, err := xmlquery.Parse(bytes.NewReader(body))
	if err != nil {
		log.Debug().Err(err).Msg("failed to parse news RSS")
		return nil, constants.ErrorMalformedBody
	}

	log.Trace().Msg("successfully parsed news RSS")

	entries := m.PluckItems(root)
	if len(entries) == 0 {
		log.Trace().Msg("no news found, exiting early")
		return nil, nil
	}

	for _, entry := range entries {
		body, err = p.client.FetchArticle(ctx, entry.Link)
		if err != nil {
			log.Debug().Err(err).Str("url", entry.Link).Msg("failed to fetch article")
			return nil, err
		}

		article, err := htmlquery.Parse(bytes.NewReader(body))
		if err != nil {
			log.Debug().Err(err).Str("url", entry.Link).Msg("failed to parse article")
			return nil, constants.ErrorMalformedBody
		}

		entry.Description = m.PickArticleDescription(article)
		entry.Content, err = m.PickArticleContents(article)
		if err != nil {
			log.Debug().Err(err).Str("url", entry.Link).Msg("error rendering article")
			return nil, constants.ErrorMalformedBody
		}
	}

	feed := m.PickSiteMeta(root)
	feed.Entries = entries
	feed.Updated = entries[0].Published

	return feed, nil
}
