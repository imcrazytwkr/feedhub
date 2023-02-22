package pixiv

import (
	"github.com/gin-gonic/gin"
	"github.com/imcrazytwkr/feedhub/providers"
	"github.com/imcrazytwkr/feedhub/routes"
)

type PixivRouter struct {
	provider providers.PixivProvider
}

func NewPixivRouter(provider providers.PixivProvider) routes.RouteContainer {
	return &PixivRouter{provider}
}

func (r *PixivRouter) Register(router gin.IRouter) {
	router.GET("/user/:id", r.handleIllustrations)
}
