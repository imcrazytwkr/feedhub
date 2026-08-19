package aws

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/providers"
	h "github.com/imcrazytwkr/feedhub/providers/aws/http"
	"github.com/imcrazytwkr/feedhub/providers/aws/mappers"
	m "github.com/imcrazytwkr/feedhub/providers/aws/models"
	"github.com/rs/zerolog"
)

type awsDirectoryProvider struct {
	client *h.AWSDirectoryClient
}

func NewAWSDirectoryProvider(httpClient *http.Client) providers.AWSDirectoryProvider {
	return &awsDirectoryProvider{h.NewAWSDirectoryClient(httpClient)}
}

func (p *awsDirectoryProvider) GetBlogs(ctx context.Context, category string, lang models.Language) (*models.Feed, error) {
	locale, supported := m.GetLocaleFor(lang)
	if !supported {
		zerolog.Ctx(ctx).Debug().Msgf("unsupported language: %q", lang)
		return nil, nil
	}
	category = strings.ToLower(strings.TrimSpace(category))

	log := zerolog.Ctx(ctx).With().Str("locale", locale.String()).Str("category", category).Logger()

	body, err := p.client.SearchBlogPosts(ctx, locale, category)
	if err != nil {
		log.Debug().Err(err).Msg("failed to fetch blog posts")
		return nil, err
	}

	log.Trace().Msg("fetched blog posts")

	var payload m.SearchResponse
	err = json.Unmarshal(body, &payload)
	if err != nil {
		log.Debug().Err(err).Msg("failed to parse blog posts")
		return nil, constants.ErrorMalformedBody
	}

	log.Trace().Msg("successfully parsed blog posts")

	entries := mappers.PluckEntries(&payload)
	if len(entries) == 0 {
		log.Trace().Msg("no blog posts found, exiting early")
		return nil, nil
	}

	feed := mappers.GenerateSiteMeta(lang, category)
	feed.Published = entries[0].Published
	feed.Entries = entries
	return feed, nil
}
