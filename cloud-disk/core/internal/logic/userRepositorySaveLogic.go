package logic

import (
	"context"
	
	"core/helper"
	"model"
	"core/internal/svc"
	"core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySaveRequest,userIdentity string) (resp *types.UserRepositorySaveResponse, err error) {
	// todo: add your logic here and delete this line
	ur := model.UserRepository{
		Identity:           helper.GetUUID(),
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                req.Ext,
		Name:               req.Name,
	}
	// 判断当前名称在该层级下是否存在
	// 插入数据库
	err=l.svcCtx.DB.Create(&ur).Error
	if err!=nil{
		return 
	}
	return &types.UserRepositorySaveResponse{
		Identity:ur.Identity,
		Ext: ur.Ext,
		Name: ur.Name,
	},nil
}
