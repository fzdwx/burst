package proxy

import (
	"context"

	"github.com/fzdwx/burst/server/internal/svc"
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

func (l *AddProxyLogic) AddProxy() error {
	// todo: add your logic here and delete this line

	return nil
}
