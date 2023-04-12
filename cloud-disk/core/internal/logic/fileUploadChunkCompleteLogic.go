package logic

import (
	"context"

	"core/helper"
	"core/internal/svc"
	"core/internal/types"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadChunkCompleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadChunkCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadChunkCompleteLogic {
	return &FileUploadChunkCompleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadChunkCompleteLogic) FileUploadChunkComplete(req *types.FileUploadChunkCompleteRequest) (resp *types.FileUploadChunkCompleteResponse, err error) {
	// todo: add your logic here and delete this line
	co := make([]cos.Object, 0)
	for _, v := range req.CosObjects{
		co = append(co,cos.Object{
			ETag: v.Etag,
			PartNumber: v.PartNumber,
		})
	}
	err=helper.CosPartUploadComplete(req.Name,req.UploadId,co)	
	if err!=nil{
		return 
	}
	return resp,nil
}