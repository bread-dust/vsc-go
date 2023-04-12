package logic

import (
	"context"
	"errors"
	"model"


	"core/internal/svc"
	"core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest,userIdentity string) (resp *types.UesrFileMoveResponse, err error) {
	// todo: add your logic here and delete this line
	// 查询用户文件夹是否存在
	up := new(model.UserRepository)
	has:=l.svcCtx.DB.Where("identity=? AND user_identity=?",req.ParentIdentity,userIdentity).First(&up).RowsAffected
	if has==0{
		return nil,errors.New("该文件夹不存在")
	}

	// 移动用户所在用户文件夹
	ur := new(model.UserRepository)
	_=l.svcCtx.DB.Model(&ur).Where("identity=?",req.Identity).Update("parent_id",up.ID)
	
	return &types.UesrFileMoveResponse{
		Result: "修改成功",
	},nil
}
