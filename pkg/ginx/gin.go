package ginx

import (
	"github.com/fzdwx/burst/pkg/logx"
	"github.com/gin-gonic/gin"
	"time"
)

func Classic() *gin.Engine {
	engine := gin.New()
	engine.Use(Logger(), gin.Recovery())
	return engine
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !logx.EnableDebug() {
			return
		}
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		if raw != "" {
			path = path + "?" + raw
		}

		log := logx.Debug().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Dur("cost", time.Now().Sub(start)).
			Int("status", c.Writer.Status()).
			Str("clientIp", c.ClientIP())

		s := c.Errors.ByType(gin.ErrorTypePrivate).String()
		if s == "" {
			s = "REQUEST"
		}
		log.Msg(s)
	}
}
