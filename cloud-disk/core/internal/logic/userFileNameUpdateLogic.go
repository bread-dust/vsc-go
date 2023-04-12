package logic

import (
	"context"
	"errors"
	"model"

	"core/internal/svc"
	"core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest,userIdentity string) (resp *types.UserFileNameUpdateResponse, err error) {
	// todo: add your logic here and delete this line
	// 判断当前名称在该层级下是否存在
	
	ur := new(model.UserRepository)
	cnt :=l.svcCtx.DB.Where("name=? AND parent_id=(select parent_id from user_repository as ur where ur.identity=?)",req.Name,req.Identity).First(&ur).RowsAffected
	if cnt >0 {
		return nil,errors.New("文件名再当前文件夹已存在")
	}

	// 修改用户文件名称
	err=l.svcCtx.DB.Model(&ur).Where("identity=? AND user_identity=?",req.Identity,userIdentity).Update("name",req.Name).Error
	if err!=nil{
		return nil,err
	}
	
	
	
	return &types.UserFileNameUpdateResponse{
		Message: "修改成功",
	},nil
}
