package logic

import (
	"context"

	"core/define"
	"core/helper"
	"core/internal/svc"
	"core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshAuthorizationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshAuthorizationLogic {
	return &RefreshAuthorizationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshAuthorizationLogic) RefreshAuthorization(req *types.RefreshAuthorizationRequest,authorization string ) (resp *types.RefreshAuthorizationResponse, err error) {
	// todo: add your logic here and delete this line
	// 解析token
	uc,err:=helper.AnalyzeToken(authorization) 
	if err!=nil{
		return nil,err
	}
	// 生成新的token
	token,err:=helper.GenerateToken(uc.Id,uc.Identity,uc.Name,define.JwtTokenExpired) 
	if err!=nil{
		return nil,err
	}

	// 生成新的refreshToken
	refreshToken,err:= helper.GenerateToken(uc.Id,uc.Identity,uc.Name,define.JwtRefreshTokenExpired)
	if err!=nil{
		return nil,err
	}
	// 返回结果
	return &types.RefreshAuthorizationResponse{
		Token:        token,
		RefreshToken: refreshToken,
	},nil
}
