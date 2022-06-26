package proxy

import (
	"net/http"

	"github.com/fzdwx/burst/server/internal/logic/proxy"
	"github.com/fzdwx/burst/server/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RemoveProxyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := proxy.NewRemoveProxyLogic(r.Context(), svcCtx)
		err := l.RemoveProxy()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
