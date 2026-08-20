package anthropic

import (
	"github.com/gin-gonic/gin"
	"github.com/imcrazytwkr/feedhub/providers"
	"github.com/imcrazytwkr/feedhub/routes"
)

type AnthropicRouter struct {
	provider providers.AnthropicProvider
}

func NewAnthropicRouter(provider providers.AnthropicProvider) routes.RouteContainer {
	return &AnthropicRouter{provider}
}

func (r *AnthropicRouter) Register(router gin.IRouter) {
	router.GET("/news", r.handleNews)
	router.GET("/engineering", r.handleEngineering)
	router.GET("/research", r.handleResearch)
	router.GET("/research/:team", r.handleResearch)
}
