package cursor

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/imcrazytwkr/feedhub/routes/common"
	"github.com/imcrazytwkr/feedhub/utils/ginutil"
)

func (r *CursorRouter) handleBlog(c *gin.Context) {
	topic, lang := common.ParseSlugWithLocale(
		strings.TrimRight(c.Param("slug"), "/"),
		strings.TrimRight(c.Param("lang"), "/"),
	)

	feed, err := r.provider.GetBlog(c.Request.Context(), topic, lang)
	if err != nil {
		ginutil.HandleError(c, err)
		return
	}

	if feed == nil {
		c.String(http.StatusNotFound, "Cursor: no blog posts found")
		return
	}

	ginutil.RenderFeed(c, feed)
}
