package ping

import (
	"github.com/fzdwx/burst/server/svc"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping test function
func Ping(svcContext *svc.ServiceContext) (string, gin.HandlerFunc) {
	return "/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	}
}