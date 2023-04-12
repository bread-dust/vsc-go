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

type FileUploadPrepareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadPrepareLogic {
	return &FileUploadPrepareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadPrepareLogic) FileUploadPrepare(req *types.FileUploadPrepareRequest) (resp *types.FileUploadPrepareResponse, err error) {
	// todo: add your logic here and delete this line
	rp := new(model.RepositoryPool)
	has:=l.svcCtx.DB.Where("hash=?",req.Md5).First(&rp).RowsAffected

	if has==0{
		return nil,errors.New("文件不存在")
	}

	if has>1{
		resp.Identity=rp.Identity
	}else{
		// 拿到fileUploadID、Key,进行分片上传
		name,uploadID,err:=helper.CosInitPart(req.Ext)
		if err!=nil{
			return nil,err
		}

		resp.Name = name
		resp.UploadId = uploadID	
	}
	return resp,nil
}
