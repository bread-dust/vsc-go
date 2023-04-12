package logic

import (
	"context"

	"ucenter/api/register/internal/svc"
	"ucenter/api/types/register"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterByphoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterByphoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterByphoneLogic {
	return &RegisterByphoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterByphoneLogic) RegisterByphone(in *register.RegReq) (*register.RegRes, error) {
	// todo: add your logic here and delete this line

	return &register.RegRes{}, nil
}
