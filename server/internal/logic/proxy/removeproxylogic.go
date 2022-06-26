package proxy

import (
	"context"

	"github.com/fzdwx/burst/server/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveProxyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveProxyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveProxyLogic {
	return &RemoveProxyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveProxyLogic) RemoveProxy() error {
	// todo: add your logic here and delete this line

	return nil
}
