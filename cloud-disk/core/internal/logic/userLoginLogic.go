package logic

import (
	"context"
	"core/define"
	"core/helper"
	"errors"
	"model"

	"core/internal/svc"
	"core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line

	// 根据用户名和密码查询用户是否存在
	user:=new(model.UserBasic)
	err = l.svcCtx.DB.Where("name=? AND password=?",req.Name,helper.Md5(req.Password)).Find(&user).Error
	if err!=nil{ // 查询过程错误
		return nil,err
	}
	if errors.Is(err,gorm.ErrRecordNotFound){ // 查询结果为空
		println(req.Name,helper.Md5(req.Password))
		return nil,errors.New("用户名或密码错误")
	}

	// 生成token
	token,err:=helper.GenerateToken(uint64(user.ID),user.Identity,user.Name,define.JwtTokenExpired)
	if err!=nil{
		return nil,err
	}

	// 刷新token
	refreshToken,err:=helper.GenerateToken(uint64(user.ID),user.Identity,user.Name,define.JwtTokenExpired)
	if err!=nil{
		return nil,err
	}

	// 返回结果
	resp = new(types.LoginResponse)
	resp.Token=token
	resp.RefreshToken = refreshToken
	return
}
