package client

import (
	"context"

	"github.com/fzdwx/burst/server/internal/svc"
	"github.com/fzdwx/burst/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConnectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConnectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConnectLogic {
	return &ConnectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConnectLogic) Connect(req *types.ClientConnectReq) (resp *types.ClientConnectResp, err error) {
	// todo: add your logic here and delete this line

	return
}
