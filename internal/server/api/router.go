package api

import (
	"github.com/fzdwx/burst/internal/server/api/ping"
	proxy2 "github.com/fzdwx/burst/internal/server/api/proxy"
	"github.com/fzdwx/burst/internal/server/api/user"
	"github.com/fzdwx/burst/internal/server/api/ws"
	"github.com/fzdwx/burst/internal/server/svc"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

func MountRouters(s *rest.Server, svcContext *svc.ServiceContext) {
	s.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/accept",
				Handler: ws.Accept(svcContext),
			},
			{
				Method:  http.MethodGet,
				Path:    "/ping",
				Handler: ping.Ping(svcContext),
			},
		},
	)

	s.AddRoutes([]rest.Route{
		{
			Method:  http.MethodGet,
			Path:    "/auth",
			Handler: user.Auth(svcContext),
		},
	},
		rest.WithPrefix("/user"),
	)

	s.AddRoutes([]rest.Route{
		{
			Method:  http.MethodPost,
			Path:    "/add/:token",
			Handler: proxy2.AddProxy(svcContext),
		},
		{
			Method:  http.MethodPost,
			Path:    "/remove/:token",
			Handler: proxy2.RemoveProxy(svcContext),
		},
	},
		rest.WithPrefix("/proxy"),
	)
}
