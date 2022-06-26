package user

import (
	"net/http"

	"github.com/fzdwx/burst/server/internal/logic/user"
	"github.com/fzdwx/burst/server/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewAuthLogic(r.Context(), svcCtx)
		resp, err := l.Auth()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
