package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"time"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)



type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func passwordMd5(password []byte) string{
		h:=md5.New()
	h.Write([]byte(password))
	h.Write(secret) //加盐
	passworddStr:=hex.EncodeToString(h.Sum(nil))
	return passworddStr
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	// 登录功能
	// 处理用户发来的请求，拿到用户名和密码
	// req.Usernmae
	// 判断用户名和密码是否和数据的一致
	user,err:=l.svcCtx.UserModel.FindOneByUsername(context.Background(),req.Username)
	if err== sqlx.ErrNotFound{
		return &types.LoginResponse{Message: "用户名不存在"},nil 	
	}
	if err!=nil{
		return &types.LoginResponse{
			logx.Errorv("findone failed")
			Message: "用户名不存在",
		},errors.New("内部错误")
	}
	if user.Password != passwordMd5([]byte(req.Password)){
		return &types.LoginResponse{Message: "用户名或密码错误"},nil
	}
	// 生成jwt
	now := time.Now().Unix()
	expire:= l.svcCtx.Config.Auth.AccessExpire
	token,err :=l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret,now,expire,user.UserId)
	if err!=nil{
		logx.Errorv("token:",err)
		return nil,errors.New("内部错误")
	}
	// 结果一致登录成功
	return &types.LoginResponse{
		Message: "登录成功",
		AccessExpire:int(now+expire),
		AccessToken:token,
		RefreshAfter:int(Now+expire/2),
	},nil
}

// jwt鉴权
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
  claims := make(jwt.MapClaims)
  claims["exp"] = iat + seconds
  claims["iat"] = iat
  claims["userId"] = userId
  token := jwt.New(jwt.SigningMethodHS256)
  token.Claims = claims
  return token.SignedString([]byte(secretKey))
}