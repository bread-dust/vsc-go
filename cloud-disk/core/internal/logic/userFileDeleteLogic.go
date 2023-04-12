package logic

import (
	"context"
	"errors"
	"model"

	"core/internal/svc"
	"core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileDeleteLogic {
	return &UserFileDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileDeleteLogic) UserFileDelete(req *types.UserFileDeleteRequest,userIdentity string) (resp *types.UserFileDeleteResponse, err error) {
	// todo: add your logic here and delete this line

	// 删除用户文件
	up := new(model.UserRepository)
	err=l.svcCtx.DB.Where("user_identity=? AND identity=?",userIdentity,req.Identity).Delete(&up).Error
	if err!=nil{
		return nil,errors.New("删错出错")
	}
	return &types.UserFileDeleteResponse{
		Message: "删除成功",
	},nil
}
