package client

import (
	"github.com/fzdwx/burst/common/errx"
	"github.com/fzdwx/burst/server/internal/logic/connect"
	"net/http"

	"github.com/fzdwx/burst/server/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ConnectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")
		if errx.CheckToken(token) {
			httpx.Error(w, errx.ErrTokenIsRequired)
			return
		}

		l := connect.NewConnectLogic(r.Context(), svcCtx)
		l.Accept(token, r, w)
	}
}
