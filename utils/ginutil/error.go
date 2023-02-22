package ginutil

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/imcrazytwkr/feedhub/models"
)

func HandleError(c *gin.Context, err error) {
	htmlErr, ok := err.(*models.HTTPError)
	if !ok {
		c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if htmlErr.StatusCode == http.StatusServiceUnavailable {
		c.Header(constants.RetryAfterHeader, strconv.Itoa(constants.DefaultTtl*60))
	}

	c.String(htmlErr.StatusCode, htmlErr.Error())
}
