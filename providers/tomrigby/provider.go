package tomrigby

import (
	"bytes"
	"context"
	"net/http"

	"github.com/antchfx/xmlquery"
	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/providers"
	h "github.com/imcrazytwkr/feedhub/providers/tomrigby/http"
	m "github.com/imcrazytwkr/feedhub/providers/tomrigby/mappers"
	"github.com/imcrazytwkr/feedhub/utils/feedutil"
	"github.com/rs/zerolog"
)

type tomRigbyProvider struct {
	client *h.TomRigbyClient
}

func NewTomRigbyProvider(httpClient *http.Client) providers.TomRigbyProvider {
	return &tomRigbyProvider{
		client: h.NewTomRigbyClient(httpClient),
	}
}

func (p *tomRigbyProvider) GetPosts(ctx context.Context) (*models.Feed, error) {
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

	entries, err := m.PluckEntries(root)
	if err != nil {
		log.Debug().Err(err).Msg("failed to pluck blog entries")
		return nil, nil
	}

	if len(entries) == 0 {
		log.Trace().Msg("no news found, exiting early")
		return nil, nil
	}

	feed := m.PickSiteMeta(root)
	feed.Entries = entries

	if feed.Published.IsZero() {
		feed.Published = feedutil.GetLastPublishedTime(entries)
	}

	if feed.Updated.IsZero() {
		feed.Updated = feedutil.GetLastUpdatedTime(entries)
	}

	return feed, nil
}
