package proxy

import (
	"context"
	"github.com/fzdwx/burst/common/errx"

	"github.com/fzdwx/burst/server/internal/svc"
	"github.com/fzdwx/burst/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProxyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddProxyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProxyLogic {
	return &AddProxyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddProxyLogic) AddProxy(req *types.AddProxyReq) (resp *types.AddProxyResp, err error) {
	// todo: add proxy
	ws := l.svcCtx.Hub.Get(req.Token)
	if ws == nil {
		return nil, errx.ErrClientNotFound
	}

	// todo
	err = ws.WriteStr("hello world")

	return
}
