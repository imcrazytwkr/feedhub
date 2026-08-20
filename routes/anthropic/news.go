package anthropic

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imcrazytwkr/feedhub/utils/ginutil"
)

func (r *AnthropicRouter) handleNews(c *gin.Context) {
	feed, err := r.provider.GetNews(c.Request.Context())
	if err != nil {
		ginutil.HandleError(c, err)
		return
	}

	if feed == nil {
		c.String(http.StatusNotFound, "Anthropic: no news found")
		return
	}

	ginutil.RenderFeed(c, feed)
}
