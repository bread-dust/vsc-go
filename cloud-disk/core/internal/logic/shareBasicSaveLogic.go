package logic

import (
	"context"
	"errors"
	"model"

	"core/helper"
	"core/internal/svc"
	"core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicSaveLogic {
	return &ShareBasicSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicSaveLogic) ShareBasicSave(req *types.ShareBasicSaveRequest,userIdentity string) (resp *types.ShareBasicSaveResponse, err error) {
	// todo: add your logic here and delete this line
	// 获取资源详细信息
	// 判断资源是否存在
	rp := new(model.RepositoryPool)
	cnt :=l.svcCtx.DB.Where("identity=?",req.RepositoryIdentity).First(&rp).RowsAffected
	if err!=nil{
		return nil,err
	}
	if cnt==0 {
		return nil,errors.New("资源不存在")
	}

	// user_repository资源保存
	ur := &model.UserRepository{
		Identity:           helper.GetUUID(),
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                rp.Ext,
		Name:               rp.Name,
	}

	err= l.svcCtx.DB.Create(&ur).Error
	if err!=nil{
		return nil,err
	}
	return &types.ShareBasicSaveResponse{
		Identity:ur.Identity,
	},nil
}
