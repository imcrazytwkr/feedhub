package tomrigby

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imcrazytwkr/feedhub/utils/ginutil"
)

func (r *TomRigbyRouter) handlePosts(c *gin.Context) {
	feed, err := r.provider.GetPosts(c.Request.Context())
	if err != nil {
		ginutil.HandleError(c, err)
		return
	}

	if feed == nil {
		c.String(http.StatusNotFound, "Thomas Rigby: no posts found")
		return
	}

	ginutil.RenderFeed(c, feed)
}
