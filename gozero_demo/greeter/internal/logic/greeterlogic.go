package logic

import (
	"context"

	"gozero_demo/greeter/internal/svc"
	"gozero_demo/greeter/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreeterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreeterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GreeterLogic {
	return &GreeterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreeterLogic) Greeter(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
