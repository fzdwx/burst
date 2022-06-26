package client

import (
	"net/http"

	"github.com/fzdwx/burst/server/internal/logic/client"
	"github.com/fzdwx/burst/server/internal/svc"
	"github.com/fzdwx/burst/server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ConnectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ClientConnectReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := client.NewConnectLogic(r.Context(), svcCtx)
		resp, err := l.Connect(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
