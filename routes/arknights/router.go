package arknights

import (
	"github.com/gin-gonic/gin"
	"github.com/imcrazytwkr/feedhub/providers"
	"github.com/imcrazytwkr/feedhub/routes"
)

type ArknightsRouter struct {
	provider providers.ArknightsProvider
}

func NewArknightsRouter(provider providers.ArknightsProvider) routes.RouteContainer {
	return &ArknightsRouter{provider}
}

func (r *ArknightsRouter) Register(router gin.IRouter) {
	router.GET("/:lang", r.handleNews)
	router.GET("", r.handleNews)
}
