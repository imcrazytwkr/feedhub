package ginutil

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/models/atom"
	"github.com/imcrazytwkr/feedhub/models/rss"
)

func RenderFeed(c *gin.Context, feed *models.Feed) {
	switch getReponseFormat(c) {
	case models.FormatAtom:
		c.Header(constants.ContentTypeHeader, atom.AtomMime)
		c.XML(http.StatusOK, atom.NewAtomFeed(feed))
	case models.FormatRss:
		c.Header(constants.ContentTypeHeader, rss.RssMime)
		c.XML(http.StatusOK, rss.NewRssFeed(feed))
	}
}

func getReponseFormat(c *gin.Context) models.Format {
	rawFormat, ok := c.Get(constants.ResponseFormatKey)
	if !ok {
		return models.FormatAtom
	}

	format, ok := rawFormat.(models.Format)
	if !ok {
		return models.FormatAtom
	}

	return format
}
