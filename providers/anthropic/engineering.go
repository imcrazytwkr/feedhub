package anthropic

import (
	"bytes"
	"context"

	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/providers/anthropic/mappers"
	"github.com/rs/zerolog"
	"golang.org/x/net/html"
)

func (p *anthropicProvider) GetEngineering(ctx context.Context) (*models.Feed, error) {
	log := zerolog.Ctx(ctx)

	body, err := p.client.FetchEngineering(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("failed to fetch engineering page")
		return nil, err
	}

	root, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		log.Debug().Err(err).Msg("failed to parse engineering page")
		return nil, constants.ErrorMalformedBody
	}

	entries := mappers.PluckEngineeringEntries(root)
	if len(entries) == 0 {
		log.Trace().Msg("no entries found, exiting early")
		return nil, nil
	}

	for _, entry := range entries {
		entry.Content, err = p.getArticleContent(ctx, entry.Link)
		if err != nil {
			return nil, err
		}
	}

	feed := mappers.EngineeringSiteMeta()
	feed.Entries = entries
	feed.Updated = entries[0].Published

	return feed, nil
}
