package ws

import (
	"github.com/fzdwx/burst"
	"github.com/fzdwx/burst/pkg/logx"
	"github.com/fzdwx/burst/pkg/protocal"
	"github.com/fzdwx/burst/pkg/result"
	"github.com/fzdwx/burst/pkg/wsx"
	"github.com/fzdwx/burst/server/api/ws/handler"
	"github.com/fzdwx/burst/server/cache"
	"github.com/fzdwx/burst/server/svc"
	"net/http"
	"time"
)

// Accept client connection
func Accept(svcContext *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// check token
		token := r.URL.Query().Get("token")
		if token == burst.EmptyStr {
			result.HttpBadRequest(w, "token not found")
			return
		}

		if !cache.ProxyInfoContainer.Has(token) {
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
		cache.ServerContainer.Put(token, ws)

		ws.MountBinaryFunc(func(bytes []byte) {
			decode, err := protocal.Decode(bytes)
			if err != nil {
				logx.Err(err).Msg("decode burst")
				return
			}

			handler.Dispatch(decode)
		})

		ws.MountCloseFunc(func(err error) {
			cache.ProxyInfoContainer.Remove(token)
			cache.ServerContainer.Remove(token)
		})

		go ws.StartReading(time.Second * 20)
		go ws.StartWriteHandler(time.Second * 5)
	}
}