package logic

import (
	"context"
	"encoding/json"
	"errors"

	"core/define"
	"core/internal/svc"
	"core/internal/types"
	"fmt"
	"model"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func mapToJson(result interface{}) string {
    // map转 json str
    jsonBytes, _ := json.Marshal(result)
    jsonStr := string(jsonBytes)
    return jsonStr
}


func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest,userIdentity string) (resp *types.UserFileListReponse, err error) {
	// todo: add your logic here and delete this line
	fmt.Printf("user_id is %d\n",req.Id)

	var ufs []*types.UserFile // 用户文件列表
	resp = new(types.UserFileListReponse) // 返回结果
	size := req.Size // 每页显示的数量
	if size == 0 {
		size = define.PageSize // 默认每页显示的数量
	}

	page := req.Page  // 当前页
	if page == 0 {
		page = 1  // 默认当前页
	}
	offset := (page-1) * size // 偏移量

	// 查询用户文件列表
	err=l.svcCtx.DB.Raw("SELECT a.id,a.identity,a.repository_identity,a.name,a.ext,b.path,b.size FROM user_repository a LEFT JOIN repository_pool b ON a.repository_identity=b.identity where a.parent_id=? AND a.user_identity=? AND a.deleted_at is null limit ?,?",req.Id,userIdentity,offset,size).Scan(&ufs).Error
	fmt.Println(mapToJson(ufs))
	if errors.Is(err,gorm.ErrRecordNotFound){ // 没有数据
		return nil,err
	}
	if err!=nil{ // 查询出错
		return nil,err
	}
	
	var ur []model.UserRepository // 用户资源列表
	// 查询用户资源总数
	count := l.svcCtx.DB.Where("parent_id =? AND user_identity=?",req.Id,userIdentity).Find(&ur).RowsAffected
	if count==0{
		return nil,err
	}
	
	
	return &types.UserFileListReponse{
		List:  ufs,
		Count: count,
	},nil
}

