package handler

import (
	"core/internal/svc"
	"core/helper"
	"core/internal/logic"
	"core/internal/types"
	"crypto/md5"
	"fmt"

	"model"
	"net/http"
	"path"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// FileUploadHandler 文件上传
func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// file是文件对象，fileHeader封装了文件基本信息
		file,fileHeader,err:= r.FormFile("file") 
		if err!=nil{
			return
		}

		// 判断文件是否已存在
		b := make([]byte,fileHeader.Size)
		_,err= file.Read(b) // 读物文件流为 b
		if err!=nil{
			return 
		}

		hash := fmt.Sprintf("%x",md5.Sum(b)) //算出文件hash值
		rp := new(model.RepositoryPool) //创建存储池对象
		// has,err:=svcCtx.DB.Where("hash=?",hash).Get(rp) //根据hash查询对象是否存在
		has:=svcCtx.DB.Where("hash=?",hash).First(&rp).RowsAffected
	
		if has>0{
			httpx.OkJson(w,&types.FileUploadResponse{Identity: rp.Identity,Ext: rp.Ext,Name: rp.Name})
			return
		}

		// 文件不存在,往COS中存储文件
		cosPath,err:=helper.CosUpload(r)
		if err!=nil{
			return 
		}

		// 往logic 传 req
		req.Name = fileHeader.Filename
		req.Ext  = path.Ext(fileHeader.Filename)
		req.Size = fileHeader.Size
		req.Hash = hash
		req.Path = cosPath

		
		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
