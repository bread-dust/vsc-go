package handler

import (
	"errors"
	"net/http"

	"core/helper"
	"core/internal/logic"
	"core/internal/svc"
	"core/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadChunkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadChunkRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		

		if r.PostForm.Get("name")==""{
			httpx.Error(w,errors.New(""))
			return
		}
		if r.PostForm.Get("upload_id")==""{
			httpx.Error(w,errors.New(""))
			return
		}
		if r.PostForm.Get("part_number")==""{
			httpx.Error(w,errors.New(""))
			return
		}

		etag,err:=helper.CosPartUpload(r)
		if err!=nil{
			httpx.Error(w,err)
			return
		}


		l := logic.NewFileUploadChunkLogic(r.Context(), svcCtx)
		resp, err := l.FileUploadChunk(&req)
		resp.Etag = etag
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
