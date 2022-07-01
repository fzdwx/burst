package api

import (
	"github.com/fzdwx/burst/server/api/ping"
	"github.com/fzdwx/burst/server/api/user"
	"github.com/fzdwx/burst/server/api/ws"
	"github.com/fzdwx/burst/server/svc"
	"github.com/gin-gonic/gin"
)

func MountRouters(e *gin.Engine, svcContext *svc.ServiceContext) {
	e.GET(ws.Accept(svcContext))
	e.GET(ping.Ping(svcContext))

	e.Group("/user").
		GET(user.Auth())
}