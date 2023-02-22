package middleware

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/rs/zerolog"
)

type responseFormat struct{}

var responseFormatKey = responseFormat{}

const rawResponseFormat = "raw_response_format"

func ResponseFormat(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		log := zerolog.Ctx(ctx)

		format, ok := ctx.Value(responseFormatKey).(models.Format)
		if ok && format != models.FormatUndefined {
			c.Request = c.Request.WithContext(log.With().Str(constants.ResponseFormatKey, format.String()).Logger().WithContext(ctx))
			c.Set(constants.ResponseFormatKey, format)
			c.Next()
			return
		}

		path := c.Request.URL.Path
		lastDot := strings.LastIndexByte(path, '.')
		if lastDot > -1 {
			stringFormat := path[lastDot+1:]
			c.Request.URL.Path = path[:lastDot]

			format = models.ParseFormat(stringFormat)
			log.Debug().
				Str(rawResponseFormat, stringFormat).
				Str(constants.ResponseFormatKey, format.String()).
				Msg("user-requested response format")
		}

		if format == models.FormatUndefined {
			c.Request.URL.Path = path

			format = models.FormatAtom
			log.Trace().
				Str(constants.ResponseFormatKey, models.FormatAtom.String()).
				Msg("no response format requested, using default")
		}

		c.Request = c.Request.WithContext(context.WithValue(ctx, responseFormatKey, format))
		engine.HandleContext(c)
	}
}
