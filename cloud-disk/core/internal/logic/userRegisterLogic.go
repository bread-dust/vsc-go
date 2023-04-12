package logic

import (
	"context"
	"errors"
	"core/helper"
	"log"
	"model"

	"core/internal/svc"
	"core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	// todo: add your logic here and delete this line
	// 判断验证码否与redis一致
	code,err := l.svcCtx.RDB.Get(l.ctx,req.Email).Result()
	if err!=nil{ // 查询redis验证码错误
		return nil,errors.New("查询redis验证码错误")
	}
	if code != req.Code{ // 验证码不对
		return nil,errors.New("验证码不对")
	}
	// 根据用户名查询用户是否已经存在
	ub := new(model.UserBasic)
	cnt := l.svcCtx.DB.Where("name=?",req.Name).Find(&ub).RowsAffected
	if cnt>0{
		return nil,errors.New("用户已存在")
	}
	user:=&model.UserBasic{
		Identity: helper.GetUUID(),
		Name:     req.Name,
		Password:  helper.Md5(req.Password),
		Email:    req.Email,	
	}
	
	// 插入用户
	err=l.svcCtx.DB.Create(user).Error
	if err!=nil{
		return nil,err
	}
	log.Println("insert user rows")

	return 
} 
