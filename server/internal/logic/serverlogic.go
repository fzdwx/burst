package logic

import (
	"context"

	"github.com/fzdwx/burst/server/internal/svc"
	"github.com/fzdwx/burst/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ServerLogic {
	return &ServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ServerLogic) Server(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
