package logic

import (
	"context"

	"github.com/fzdwx/burst/internal/svc"
	"github.com/fzdwx/burst/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BurstLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBurstLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BurstLogic {
	return &BurstLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BurstLogic) Burst(req *types.Request) (resp *types.Response, err error) {
	return &types.Response{Message: "hello " + req.Name}, nil
}
