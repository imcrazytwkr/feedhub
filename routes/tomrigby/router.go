package tomrigby

import (
	"github.com/gin-gonic/gin"
	"github.com/imcrazytwkr/feedhub/providers"
	"github.com/imcrazytwkr/feedhub/routes"
)

type TomRigbyRouter struct {
	provider providers.TomRigbyProvider
}

func NewTomRigbyRouter(provider providers.TomRigbyProvider) routes.RouteContainer {
	return &TomRigbyRouter{provider}
}

func (r *TomRigbyRouter) Register(router gin.IRouter) {
	router.GET("", r.handlePosts)
}
