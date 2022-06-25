package handler

import (
	"net/http"

	"github.com/fzdwx/burst/internal/logic"
	"github.com/fzdwx/burst/internal/svc"
	"github.com/fzdwx/burst/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func BurstHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewBurstLogic(r.Context(), svcCtx)
		resp, err := l.Burst(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
