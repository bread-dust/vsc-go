package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"api/internal/svc"
	"api/internal/types"
	"model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var secret = []byte("dengliwei")
type SignupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignupLogic {
	return &SignupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignupLogic) Signup(req *types.SignupRequest) (resp *types.SignupResponse, err error) {
	// todo: add your logic here and delete this line
	// 业务逻辑
	fmt.Printf("req:%v\n",req)
	// 用户注册
	
	// 参数校验
	if req.RePassword!=req.Password{
		return nil,errors.New("两次输入的密码不一致")
	}
	// 查询user是否存在
	logx.Infov(req) //json.Marshal(req)
	u, err:= l.svcCtx.UserModel.FindOneByUsername(l.ctx,req.Username)
	// 查询数据库失败1
	if err!=nil&&err!=sqlx.ErrNotFound{
		logx.Errorf("findonebyusername",err)
		return nil,errors.New("内部错误")
	}
	// 查到记录，用户名已经注册
	if u!=nil{
		return nil,errors.New("用户名已存在 ")
	}

	// 没查到记录
	// 查到记录
	// 生成userID
	// 加密密码(md5，加盐)
	h:=md5.New()
	h.Write([]byte(req.Password))
	h.Write(secret) //加盐
	passworddStr:=hex.EncodeToString(h.Sum(nil))

	user:=&model.User{
		UserId:     time.Now().Unix(), // 雪花算法生成
		Username:   req.Username,
		Password:   passworddStr,
		Gender:     int64(req.Gender),
		CreateTime: time.Time{},
		UpdateTime: time.Time{},
	}
	if _,err:=l.svcCtx.UserModel.Insert(context.Background(),user);err!=nil{
		return nil,err
	}


	return &types.SignupResponse{Message:"success"},nil
	
}
