package logic

import (
	"context"

	"github.com/fzdwx/burst/client/internal/svc"
	"github.com/fzdwx/burst/client/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClientLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientLogic {
	return &ClientLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClientLogic) Client(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
