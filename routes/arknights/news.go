package arknights

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/imcrazytwkr/feedhub/utils/ginutil"
)

func (r *ArknightsRouter) handleNews(c *gin.Context) {
	lang := c.Param("lang")
	if len(lang) == 0 {
		lang = "en"
	} else {
		lang = strings.TrimRight(lang, "/")
	}

	feed, err := r.provider.GetNews(c.Request.Context(), lang)
	if err != nil {
		ginutil.HandleError(c, err)
		return
	}

	if feed == nil {
		c.String(http.StatusNotFound, "Arknights: no news found")
		return
	}

	ginutil.RenderFeed(c, feed)
}
