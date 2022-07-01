package user

import (
	"github.com/fzdwx/burst/server/cache"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"net/http"
)

// Auth The current role is to generate tokens
func Auth() (string, gin.HandlerFunc) {
	return "auth", func(context *gin.Context) {
		token := xid.New().String()

		cache.ProxyInfoContainer.Put(token)
		context.String(http.StatusOK, token)
	}
}