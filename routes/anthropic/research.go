package anthropic

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/imcrazytwkr/feedhub/utils/ginutil"
)

func (r *AnthropicRouter) handleResearch(c *gin.Context) {
	team := strings.Trim(c.Param("team"), "/")

	feed, err := r.provider.GetResearch(c.Request.Context(), team)
	if err != nil {
		ginutil.HandleError(c, err)
		return
	}

	if feed == nil {
		c.String(http.StatusNotFound, "Anthropic: no research found")
		return
	}

	ginutil.RenderFeed(c, feed)
}
