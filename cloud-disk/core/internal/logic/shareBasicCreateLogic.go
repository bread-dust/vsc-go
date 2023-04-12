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

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest,userIdentity string) (resp *types.ShareBasicCreateResponse, err error) {
	// todo: add your logic here and delete this line
	// 判断用户文件是否存在
	ur := new(model.UserRepository)
	// has,err:=l.svcCtx.Engine.Where("identity=?",req.UserRepositoryIdentity).Get(ur)
	cnt := l.svcCtx.DB.Where("identity=?",req.UserRepositoryIdentity).First(&ur).RowsAffected
	if err!=nil{
		return nil,err
	}
	if cnt==0{
		return nil,errors.New("数据不存在")
	}

	// 创建分享
	data:=&model.ShareBasic{
		Identity: helper.GetUUID(),
		UserIdentity: userIdentity,
		RepositoryIdentity: ur.RepositoryIdentity,
		UserRepositoryIdentity: req.UserRepositoryIdentity,
		ExpiredTime: req.ExpiredTime,
	}
	err=l.svcCtx.DB.Create(&data).Error
	if err!=nil{
		return nil,err
	}

	return &types.ShareBasicCreateResponse{
		Identity: data.Identity,
	},nil
}
