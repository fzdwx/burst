package ws

import (
	cache2 "github.com/fzdwx/burst/internal/cache"
	"github.com/fzdwx/burst/internal/logx"
	"github.com/fzdwx/burst/internal/protocal"
	"github.com/fzdwx/burst/internal/result"
	"github.com/fzdwx/burst/internal/server/api/ws/handler"
	"github.com/fzdwx/burst/internal/server/svc"
	"github.com/fzdwx/burst/internal/wsx"
	"net/http"
	"time"
)

// Accept client connection
func Accept(svcContext *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// check token
		token := r.URL.Query().Get("token")
		if token == internal.EmptyStr {
			result.HttpBadRequest(w, "token not found")
			return
		}

		if !cache2.ProxyInfoContainer.Has(token) {
			result.HttpBadRequest(w, "token is not valid")
			return
		}

		// upgrade to websocket
		conn, err := svcContext.WsUpgrader.Upgrade(w, r, nil)
		if err != nil {
			result.HttpBadRequest(w, "upgrade to websocket fail:"+err.Error())
			return
		}

		// init websocket client
		ws := wsx.NewClassicWsx(conn)
		cache2.ServerContainer.Put(token, ws)

		ws.MountBinaryFunc(func(bytes []byte) {
			decode, err := protocal.Decode(bytes)
			if err != nil {
				logx.Err(err).Msg("decode burst")
				return
			}

			handler.Dispatch(decode)
		})

		ws.MountCloseFunc(func(err error) {
			cache2.ProxyInfoContainer.Remove(token)
			cache2.ServerContainer.Remove(token)
		})

		go ws.StartReading(0)
		go ws.StartWriteHandler(time.Second * 5)
	}
}
