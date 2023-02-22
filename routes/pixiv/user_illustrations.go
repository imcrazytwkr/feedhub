package pixiv

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imcrazytwkr/feedhub/utils/ginutil"
	"github.com/rs/zerolog"
)

func (r *PixivRouter) handleIllustrations(c *gin.Context) {
	log := zerolog.Ctx(c.Request.Context()).With().
		Str("service", "pixiv").
		Str("handler", "user_illustrations").
		Logger()

	rawUserId := c.Param("id")
	if len(rawUserId) == 0 {
		log.Debug().Msg("user_id param not provided")
		c.Status(http.StatusNotFound)
		return
	}

	userId, err := strconv.Atoi(rawUserId)
	if err != nil || userId < 1 {
		log.Debug().Str("raw_user_id", rawUserId).Err(err).Msg("invalid user_id")
		c.Status(http.StatusNotFound)
		return
	}

	ctx := log.With().Int("user_id", userId).Logger().WithContext(c.Request.Context())

	feed, err := r.provider.GetUserIllustrations(ctx, userId)
	if err != nil {
		ginutil.HandleError(c, err)
		return
	}

	if feed == nil {
		c.String(http.StatusNotFound, "Pixiv: user #%d not found", userId)
		return
	}

	ginutil.RenderFeed(c, feed)
}
