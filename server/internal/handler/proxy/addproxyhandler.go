package proxy

import (
	"net/http"

	"github.com/fzdwx/burst/server/internal/logic/proxy"
	"github.com/fzdwx/burst/server/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddProxyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := proxy.NewAddProxyLogic(r.Context(), svcCtx)
		err := l.AddProxy()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
