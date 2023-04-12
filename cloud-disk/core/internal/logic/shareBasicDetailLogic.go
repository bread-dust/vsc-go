package logic

import (
	"context"

	"core/internal/svc"
	"core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

)

type ShareBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(req *types.ShareBasicDetailRequest) (resp *types.ShareBasicDetailResponse, err error) {
	// todo: add your logic here and delete this line
	// 对分享记录的点击次数进行+1
	// _,err=l.svcCtx.Engine.Exec("UPDATE share_basic SET click_num = click_num +1 WHERE identity=?",req.Identity)
	err=l.svcCtx.DB.Exec("UPDATE share_basic SET click_num = click_num +1 WHERE identity=?",req.Identity).Error
	if err!=nil{
		return nil,err
	}
	// 获取资源的详细信息
	resp = new(types.ShareBasicDetailResponse)
	err = l.svcCtx.DB.Raw("SELECT share_basic.repository_identity,user_repository.name,repository_pool.ext,repository_pool.size,repository_pool.path FROM share_basic LEFT JOIN repository_pool ON share_basic.repository_identity=repository_pool.identity LEFT JOIN user_repository ON user_repository.identity=share_basic.user_repository_identity WHERE share_basic.identity=?",req.Identity).Scan(&resp).Error
	if err!=nil{
		return nil,err
	}
	return 
}
