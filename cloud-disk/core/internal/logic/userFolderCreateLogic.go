package logic

import (
	"context"
	"errors"
	"core/helper"
	"model"

	"core/internal/svc"
	"core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateRequest,userIdentity string) (resp *types.UserFolderCreateResponse, err error) {
	// todo: add your logic here and delete this line
	// 判断当前名称在该层级下是否存在
	ur:=new(model.UserRepository)
	cnt:=l.svcCtx.DB.Where("name=? AND parent_id=?",req.Name,req.ParentId).First(&ur).RowsAffected

	if cnt >0 {
		return nil,errors.New("文件名再当前文件夹已存在")
	}

	// 创建文件夹
	data:=&model.UserRepository{
		Identity:           helper.GetUUID(),
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		Name:               req.Name,
	}

	err=l.svcCtx.DB.Create(&data).Error

	if err!=nil{
		return nil,errors.New("文件夹创建失败")
	}
	return &types.UserFolderCreateResponse{
		Identity: data.Identity,
	},nil
}
