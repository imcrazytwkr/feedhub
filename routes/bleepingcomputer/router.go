package bleepingcomputer

import (
	"github.com/gin-gonic/gin"
	"github.com/imcrazytwkr/feedhub/providers"
	"github.com/imcrazytwkr/feedhub/routes"
)

type BleepingComputerRouter struct {
	provider providers.BleepingComputerProvider
}

func NewBleepingComputerRouter(provider providers.BleepingComputerProvider) routes.RouteContainer {
	return &BleepingComputerRouter{provider}
}

func (r *BleepingComputerRouter) Register(router gin.IRouter) {
	router.GET("", r.handleNews)
}
