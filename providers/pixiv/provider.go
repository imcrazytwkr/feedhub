package pixiv

import (
	"context"
	"net/http"

	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/providers"
	h "github.com/imcrazytwkr/feedhub/providers/pixiv/http"
	m "github.com/imcrazytwkr/feedhub/providers/pixiv/mappers"
	"github.com/imcrazytwkr/feedhub/utils/logutil"
	"github.com/rs/zerolog"
	"github.com/valyala/fastjson"
)

type pixivProvider struct {
	parserPool *fastjson.ParserPool
	client     *h.PixivClient
}

func NewPixivProvider(parserPool *fastjson.ParserPool, httpClient *http.Client) providers.PixivProvider {
	return &pixivProvider{
		parserPool: parserPool,
		client:     h.NewPixivClient(httpClient),
	}
}

func (p *pixivProvider) GetUserIllustrations(ctx context.Context, userId int) (*models.Feed, error) {
	log := zerolog.Ctx(ctx)

	// Latest illustration IDs
	body, err := p.client.FetchUserIllustrations(ctx, userId)
	if err != nil {
		log.Debug().Err(err).Msg("failed to fetch illustration ids")
		return nil, err
	}

	log.Trace().Msg("fetched illustration ids")

	parser := p.parserPool.Get()
	defer p.parserPool.Put(parser)

	payload, err := parser.ParseBytes(body)
	if err != nil {
		log.Debug().Err(err).Msg("failed to parse illustration ids body")
		return nil, constants.ErrorMalformedBody
	}

	log.Trace().Msg("successfully parsed illustration ids body")

	illustIds, err := m.PluckIllustrationIds(payload)
	if err != nil {
		log.Debug().Err(err).Msg("failed to extract illustration ids")
		return nil, models.NewHttpError(http.StatusBadGateway, err)
	}

	if len(illustIds) == 0 {
		log.Trace().Msg("no illustrations found, exiting early")
		return nil, nil
	}

	log.Trace().Array("illust_ids", logutil.IntArray(illustIds)).Msg("parsed illustration ids")

	/**
	 * @TODO: consider only fetching last N illustrations for performance reasons
	 */

	// Detailed illustration data
	body, err = p.client.FetchUserIllustrationsData(ctx, userId, illustIds)
	if err != nil {
		log.Debug().Err(err).Msg("failed to fetch illustration data")
		return nil, err
	}

	log.Trace().Msg("successfully fetched illustration data")

	payload, err = parser.ParseBytes(body)
	if err != nil {
		log.Debug().Err(err).Msg("failed to parse illustration data body")
		return nil, constants.ErrorMalformedBody
	}

	log.Trace().Msg("successfully parsed illustration data body")

	illustrations, err := m.PluckIllustrationEntries(payload)
	if err != nil {
		log.Debug().Err(err).Msg("failed to extract illustration data")
		return nil, err
	}

	if len(illustrations) == 0 {
		log.Trace().Msg("no illustration meta found, exiting early")
		return nil, nil
	}

	feed := m.ExtractFeedFromUserMeta(payload)
	if feed == nil {
		log.Debug().Msg("could not parse user meta")
		return nil, constants.ErrorMalformedBody
	}

	feed.Entries = illustrations
	feed.Author = illustrations[0].Author
	feed.Updated = illustrations[0].Updated
	feed.Published = illustrations[0].Published

	return feed, nil
}
