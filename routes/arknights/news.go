package arknights

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/utils/ginutil"
)

func (r *ArknightsRouter) handleNews(c *gin.Context) {
	lang := models.ParseLanguage(strings.TrimRight(c.Param("lang"), "/"))
	if lang == models.LanguageUnknown {
		lang = models.LanguageEn
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
