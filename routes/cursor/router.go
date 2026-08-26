package cursor

import (
	"github.com/gin-gonic/gin"
	"github.com/imcrazytwkr/feedhub/providers"
	"github.com/imcrazytwkr/feedhub/routes"
)

type CursorRouter struct {
	provider providers.CursorProvider
}

func NewCursorRouter(provider providers.CursorProvider) routes.RouteContainer {
	return &CursorRouter{provider}
}

func (r *CursorRouter) Register(router gin.IRouter) {
	blogs := router.Group("/blog")
	blogs.GET("", r.handleBlog)
	blogs.GET("/:slug", r.handleBlog)
	blogs.GET("/:slug/:lang", r.handleBlog)
}
