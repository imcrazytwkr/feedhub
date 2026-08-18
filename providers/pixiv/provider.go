package pixiv

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/providers"
	h "github.com/imcrazytwkr/feedhub/providers/pixiv/http"
	"github.com/imcrazytwkr/feedhub/providers/pixiv/mappers"
	m "github.com/imcrazytwkr/feedhub/providers/pixiv/models"
	"github.com/imcrazytwkr/feedhub/utils/feedutil"
	"github.com/imcrazytwkr/feedhub/utils/logutil"
	"github.com/rs/zerolog"
)

type pixivProvider struct {
	client *h.PixivClient
}

func NewPixivProvider(httpClient *http.Client) providers.PixivProvider {
	return &pixivProvider{
		client: h.NewPixivClient(httpClient),
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

	var idsPayload m.Response[m.IllustrationIDsBody]
	err = json.Unmarshal(body, &idsPayload)
	if err != nil {
		log.Debug().Err(err).Msg("failed to parse illustration ids body")
		return nil, constants.ErrorMalformedBody
	}

	log.Trace().Msg("successfully parsed illustration ids body")

	illustIds, err := mappers.PluckIllustrationIds(&idsPayload)
	if err != nil {
		log.Debug().Err(err).Msg("failed to extract illustration ids")
		return nil, models.NewHttpError(http.StatusBadGateway, err)
	}

	if len(illustIds) == 0 {
		log.Trace().Msg("no illustrations found, exiting early")
		return nil, nil
	}

	log.Trace().
		Array("illust_ids", logutil.IntArray(illustIds)).
		Msgf("parsed %d illustration ids", len(illustIds))

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

	var dataPayload m.Response[m.IllustrationDataBody]
	err = json.Unmarshal(body, &dataPayload)
	if err != nil {
		log.Debug().Err(err).Msg("failed to parse illustration data body")
		return nil, constants.ErrorMalformedBody
	}

	log.Trace().Msg("successfully parsed illustration data body")

	illustrations, err := mappers.PluckIllustrationEntries(&dataPayload)
	if err != nil {
		log.Debug().Err(err).Msg("failed to extract illustration data")
		return nil, err
	}

	if len(illustrations) == 0 {
		log.Trace().Msg("no illustration meta found, exiting early")
		return nil, nil
	}

	log.Trace().Msgf("found %d illustrations", len(illustrations))

	feed := mappers.ExtractFeedFromUserMeta(&dataPayload.Body.ExtraData.Meta)
	if feed == nil {
		log.Debug().Msg("could not parse user meta")
		return nil, constants.ErrorMalformedBody
	}

	feed.Entries = illustrations
	feed.Author = illustrations[0].Author
	feed.Updated = feedutil.GetLastUpdatedTime(illustrations)
	feed.Published = feedutil.GetLastPublishedTime(illustrations)

	return feed, nil
}
