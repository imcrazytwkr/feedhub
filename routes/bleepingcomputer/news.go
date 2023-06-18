package bleepingcomputer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imcrazytwkr/feedhub/utils/ginutil"
)

func (r *BleepingComputerRouter) handleNews(c *gin.Context) {
	feed, err := r.provider.GetNews(c.Request.Context())
	if err != nil {
		ginutil.HandleError(c, err)
		return
	}

	if feed == nil {
		c.String(http.StatusNotFound, "BleepingComputer: no news found")
		return
	}

	ginutil.RenderFeed(c, feed)
}
