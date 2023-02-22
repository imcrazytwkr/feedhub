package constants

import (
	"net/http"

	"github.com/imcrazytwkr/feedhub/models"
)

var ErrorMalformedBody = models.NewHttpError(http.StatusBadGateway, "server responded with malformed body")
