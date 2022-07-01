package ws

import (
	"fmt"
	"github.com/fzdwx/burst"
	"github.com/fzdwx/burst/pkg/ginx"
	"github.com/fzdwx/burst/pkg/wsx"
	"github.com/fzdwx/burst/server/cache"
	"github.com/fzdwx/burst/server/svc"
	"github.com/gin-gonic/gin"
	"time"
)

// Accept  todo save client to cache
func Accept(svcContext *svc.ServiceContext) (string, gin.HandlerFunc) {
	return "/accept", func(context *gin.Context) {
		// check token
		token := context.Query("token")
		if token == burst.EmptyStr {
			ginx.Error("token not found", context)
			return
		}

		if !cache.ProxyInfoContainer.Has(token) {
			ginx.Error("token is not valid", context)
			return
		}

		// upgrade to websocket
		conn, err := svcContext.WsUpgrader.Upgrade(context.Writer, context.Request, nil)

		if err != nil {
			ginx.Error("upgrade to websocket fail", context)
			return
		}

		ws := wsx.NewClassicWsx(conn)

		ws.MountTextFunc(func(text string) {
			fmt.Println(text)
			ws.WriteText("我草11111")
		})

		ws.MountCloseFunc(func(err error) {
			cache.ProxyInfoContainer.Remove(token)
		})

		go ws.StartReading(time.Second * 20)
		go ws.StartWriteHandler(time.Second * 5)
		ws.WriteText("hello world")
	}
}