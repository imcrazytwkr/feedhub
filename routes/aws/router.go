package aws

import (
	"github.com/gin-gonic/gin"
	"github.com/imcrazytwkr/feedhub/providers"
	"github.com/imcrazytwkr/feedhub/routes"
)

type AWSRouter struct {
	provider providers.AWSDirectoryProvider
}

func NewAWSRouter(provider providers.AWSDirectoryProvider) routes.RouteContainer {
	return &AWSRouter{provider}
}

func (r *AWSRouter) Register(router gin.IRouter) {
	blogs := router.Group("/blogs")
	blogs.GET("", r.handleBlogs)
	blogs.GET("/:slug", r.handleBlogs)
	blogs.GET("/:slug/:lang", r.handleBlogs)
}
