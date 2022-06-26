package proxy

import (
	"net/http"

	"github.com/fzdwx/burst/server/internal/logic/proxy"
	"github.com/fzdwx/burst/server/internal/svc"
	"github.com/fzdwx/burst/server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddProxyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddProxyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := proxy.NewAddProxyLogic(r.Context(), svcCtx)
		resp, err := l.AddProxy(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
