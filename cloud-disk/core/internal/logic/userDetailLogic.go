package logic

import (
	"context"
	"errors"
	"fmt"

	"core/internal/svc"
	"core/internal/types"
	"model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailResponse, err error) {
	// todo: add your logic here and delete this line
	 
	fmt.Printf("user_identity is %s\n",req.Identity)
	resp = &types.UserDetailResponse{} // 返回结果
	// 判断用户是否存在
	ub := new(model.UserBasic)
	has:=l.svcCtx.DB.Where("identity=?",req.Identity).First(&ub).RowsAffected
	if err!=nil{
		return nil,err
	}
	if has==0{
		return nil,errors.New("user not found")
	}
	resp.Name = ub.Name
	resp.Email = ub.Email

	return
}

