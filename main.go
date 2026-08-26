package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imcrazytwkr/feedhub/middleware"
	anp "github.com/imcrazytwkr/feedhub/providers/anthropic"
	akp "github.com/imcrazytwkr/feedhub/providers/arknights"
	awsp "github.com/imcrazytwkr/feedhub/providers/aws"
	bp "github.com/imcrazytwkr/feedhub/providers/bleepingcomputer"
	cp "github.com/imcrazytwkr/feedhub/providers/cursor"
	pp "github.com/imcrazytwkr/feedhub/providers/pixiv"
	anr "github.com/imcrazytwkr/feedhub/routes/anthropic"
	akr "github.com/imcrazytwkr/feedhub/routes/arknights"
	awsr "github.com/imcrazytwkr/feedhub/routes/aws"
	br "github.com/imcrazytwkr/feedhub/routes/bleepingcomputer"
	cr "github.com/imcrazytwkr/feedhub/routes/cursor"
	pr "github.com/imcrazytwkr/feedhub/routes/pixiv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	if gin.IsDebugging() {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	listenHost, err := getHost()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	listenPort, err := getPort()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	engine := gin.New()
	engine.Use(middleware.DefaultLogger(), gin.Recovery())
	engine.Use(middleware.RequestId())
	engine.Use(middleware.ResponseFormat(engine))

	anthropicProvider := anp.NewAnthropicProvider(http.DefaultClient)
	anr.NewAnthropicRouter(anthropicProvider).Register(engine.Group("/anthropic"))

	arknightsProvider := akp.NewArknightsProvider(http.DefaultClient)
	akr.NewArknightsRouter(arknightsProvider).Register(engine.Group("/arknights"))

	awsProvider := awsp.NewAWSDirectoryProvider(http.DefaultClient)
	awsr.NewAWSRouter(awsProvider).Register(engine.Group("/aws"))

	bleepingComputerProvider := bp.NewBleepingComputerProvider(http.DefaultClient)
	br.NewBleepingComputerRouter(bleepingComputerProvider).Register(engine.Group("/bleepingcomputer"))

	cursorProvider := cp.NewCursorProvider(http.DefaultClient)
	cr.NewCursorRouter(cursorProvider).Register(engine.Group("/cursor"))

	pixivProvider := pp.NewPixivProvider(http.DefaultClient)
	pr.NewPixivRouter(pixivProvider).Register(engine.Group("/pixiv"))

	engine.Run(net.JoinHostPort(listenHost, listenPort))
}

func getHost() (string, error) {
	listenHost := os.Getenv("HOST")
	if len(listenHost) == 0 {
		return "", nil
	}

	// ParseIP returns nil on invalid IP
	if net.ParseIP(listenHost) == nil {
		return "", fmt.Errorf("listen host %q is not a valid IP address", listenHost)
	}

	return listenHost, nil
}

func getPort() (string, error) {
	listenPort := os.Getenv("PORT")
	if len(listenPort) == 0 {
		log.Debug().Msg("Listen port is unset or empty, falling back to default")
		listenPort = "8080"
	}

	// Checking is port number fits in Uint16
	_, err := strconv.ParseUint(listenPort, 10, 16)
	if err != nil {
		return "", fmt.Errorf("listen port number %q is invalid; Valid range is 0-65535", listenPort)
	}

	return listenPort, nil
}
