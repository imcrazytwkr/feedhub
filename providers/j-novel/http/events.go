package http

/* curl 'https://labs.j-novel.club/app/v1/events?limit=10&sort=launch&start_date=2023-06-22T18%3A05%3A00.000Z&end_date='
await fetch("https://labs.j-novel.club/app/v1/events?limit=10&sort=launch&start_date=2023-06-22T18%3A05%3A00.000Z&end_date=", {
    "credentials": "omit",
    "headers": {
        "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/114.0",
        "Accept": "*\/*",
        "Accept-Language": "en-US,en;q=0.5",
        "Sec-Fetch-Dest": "empty",
        "Sec-Fetch-Mode": "cors",
        "Sec-Fetch-Site": "same-site",
        "Sec-GPC": "1",
    },
    "referrer": "https://j-novel.club/",
    "method": "GET",
    "mode": "cors"
});
*/

import (
	"context"
	"net/http"

	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/utils/httputil"
	"github.com/rs/zerolog"
)

func (b *JNovelClient) FetchCalendarEvents(ctx context.Context) ([]byte, error) {
	log := zerolog.Ctx(ctx)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, calendarEventsUrl, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to create request to fetch news feed")
		return nil, models.NewHttpError(http.StatusInternalServerError, nil)
	}

	req.Header = http.Header{
		constants.UserAgent:            {headerUserAgent},
		constants.AcceptHeader:         {headerAccept},
		constants.AcceptLanguageHeader: {headerAcceptLanguage},
		constants.OriginHeader:         {headerOrigin},
		constants.ReferrerHeader:       {headerReferrer},
	}

	return httputil.FetchRequest(b.httpClient, req)
}
