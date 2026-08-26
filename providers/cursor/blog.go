package cursor

import (
	"bytes"
	"context"
	"strings"

	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/providers/cursor/mappers"
	m "github.com/imcrazytwkr/feedhub/providers/cursor/models"
	"github.com/rs/zerolog"
	"golang.org/x/net/html"
)

func (p *cursorProvider) GetBlog(ctx context.Context, topic string, lang models.Language) (*models.Feed, error) {
	locale, supported := m.GetLocaleFor(lang)
	if !supported {
		zerolog.Ctx(ctx).Debug().Msgf("unsupported language: %q", lang)
		return nil, nil
	}

	topic = strings.ToLower(strings.TrimSpace(topic))
	log := zerolog.Ctx(ctx).With().Str("locale", locale.String()).Str("topic", topic).Logger()

	body, err := p.client.FetchBlog(ctx, locale, topic)
	if err != nil {
		log.Debug().Err(err).Msg("failed to fetch blog listing")
		return nil, err
	}

	root, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		log.Debug().Err(err).Msg("failed to parse blog listing")
		return nil, constants.ErrorMalformedBody
	}

	entries := mappers.PluckBlogEntries(root)
	if len(entries) == 0 {
		log.Trace().Msg("no entries found, exiting early")
		return nil, nil
	}

	for _, entry := range entries {
		entry.Author, entry.Content, err = p.getArticleData(ctx, entry.Link)
		if err != nil {
			return nil, err
		}
	}

	feed := mappers.BlogSiteMeta(topic, lang, mappers.PluckBlogDescription(root))
	feed.Entries = entries
	feed.Updated = entries[0].Published

	return feed, nil
}
