package anthropic

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imcrazytwkr/feedhub/utils/ginutil"
)

func (r *AnthropicRouter) handleEngineering(c *gin.Context) {
	feed, err := r.provider.GetEngineering(c.Request.Context())
	if err != nil {
		ginutil.HandleError(c, err)
		return
	}

	if feed == nil {
		c.String(http.StatusNotFound, "Anthropic: no engineering posts found")
		return
	}

	ginutil.RenderFeed(c, feed)
}
