package ws

import (
	"github.com/fzdwx/burst/pkg/logx"
	"github.com/fzdwx/burst/pkg/wsx"
	"github.com/fzdwx/burst/server/svc"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Accept(svcContext *svc.ServiceContext) (string, gin.HandlerFunc) {
	return "/accept", func(context *gin.Context) {
		conn, err := svcContext.WsUpgrader.Upgrade(context.Writer, context.Request, nil)

		if err != nil {
			logx.Err(err).Msg("upgrade to websocket fail")
			http.Error(context.Writer, "upgrade to websocket fail", http.StatusInternalServerError)
			return
		}

		ws := wsx.NewClassicWsx(conn)
		go ws.StartReading(time.Second * 20)
		go ws.StartWriteHandler(time.Second * 5)
		ws.WriteText("hello world")
	}
}