package handler

import (
	"github.com/fzdwx/burst/server/internal/handler/client"
	"github.com/fzdwx/burst/server/internal/svc"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

func RegisterCustomHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/connect",
				Handler: client.ConnectHandler(serverCtx),
			},
		},
	)
}
