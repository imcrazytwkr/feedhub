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

func (p *anthropicProvider) GetResearch(ctx context.Context, team string) (*models.Feed, error) {
	log := zerolog.Ctx(ctx)

	body, err := p.client.FetchResearch(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("failed to fetch research page")
		return nil, err
	}

	root, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		log.Debug().Err(err).Msg("failed to parse research page")
		return nil, constants.ErrorMalformedBody
	}

	entries := mappers.PluckResearchEntries(root, team)
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

	feed := mappers.ResearchSiteMeta(team)
	feed.Entries = entries
	feed.Updated = entries[0].Published

	return feed, nil
}
