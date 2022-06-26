package client

import (
	"errors"
	"net/http"

	"github.com/fzdwx/burst/server/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var (
	tokenReq = errors.New("token is required")
)

func ConnectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")
		if token == "" {
			httpx.Error(w, tokenReq)
			return
		}

		conn, err := svcCtx.Hub.UpgradeToWs(w, r)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		client := svcCtx.Hub.AddClient(conn)

		//client.OnBinary(func(data []byte) {
		//
		//})
		//client.OnText(func(data string) {
		//
		//})

		go client.WritePump()
		go client.ReadPump()
	}
}
