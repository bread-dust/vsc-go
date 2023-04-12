package logic

import (
	"context"
	"model"

	"core/internal/svc"
	"core/internal/types"
	"core/helper"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// FileUpload 文件上传
func (l *FileUploadLogic) FileUpload(req *types.FileUploadRequest) (resp *types.FileUploadResponse, err error) {
	// todo: add your logic here and delete this line
	rp := &model.RepositoryPool{ // 生成资源池条目
		Identity: helper.GetUUID(), // 生成UUID
		Hash:     req.Hash,
		Name:     req.Name,
		Ext:      req.Ext,
		Size:     req.Size,
		Path:     req.Path,
	}
	err =  l.svcCtx.DB.Create(&rp).Error // 插入数据库
	if err!=nil{
		return nil,err
	}

	resp = new(types.FileUploadResponse) // 返回结果
	resp.Identity = rp.Identity
	resp.Ext = rp.Ext
	resp.Name = rp.Name
	
	return 

}
