package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/imcrazytwkr/feedhub/middleware"
	pp "github.com/imcrazytwkr/feedhub/providers/pixiv"
	pr "github.com/imcrazytwkr/feedhub/routes/pixiv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fastjson"
)

func main() {
	if gin.IsDebugging() {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	engine := gin.New()
	engine.Use(middleware.DefaultLogger(), gin.Recovery())
	engine.Use(middleware.ResponseFormat(engine))

	parserPool := &fastjson.ParserPool{}

	pixivProvider := pp.NewPixivProvider(parserPool, http.DefaultClient)
	pr.NewPixivRouter(pixivProvider).Register(engine.Group("/pixiv"))
	engine.Run("127.0.0.1:5000")
}
