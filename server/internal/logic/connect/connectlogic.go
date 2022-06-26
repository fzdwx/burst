package connect

import (
	"context"
	"github.com/fzdwx/burst/server/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConnectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Logic {
	return &Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Logic) Accept(token string, r *http.Request, w http.ResponseWriter) {
	conn, err := l.svcCtx.Hub.UpgradeToWs(w, r)
	if err != nil {
		httpx.Error(w, err)
		return
	}

	// TODO check token is valid

	client := l.svcCtx.Hub.AddClient(conn, token)

	client.OnBinary(func(data []byte) {

	})

	client.OnText(func(data string) {

	})

	go client.React()

}
