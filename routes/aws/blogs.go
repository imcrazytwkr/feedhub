package aws

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/imcrazytwkr/feedhub/routes/common"
	"github.com/imcrazytwkr/feedhub/utils/ginutil"
)

func (r *AWSRouter) handleBlogs(c *gin.Context) {
	category, lang := common.ParseSlugWithLocale(
		strings.TrimRight(c.Param("slug"), "/"),
		strings.TrimRight(c.Param("lang"), "/"),
	)

	feed, err := r.provider.GetBlogs(c.Request.Context(), category, lang)
	if err != nil {
		ginutil.HandleError(c, err)
		return
	}

	if feed == nil {
		c.String(http.StatusNotFound, "AWS: no blog posts found")
		return
	}

	ginutil.RenderFeed(c, feed)
}
