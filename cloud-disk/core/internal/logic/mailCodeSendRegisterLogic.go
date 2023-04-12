package logic

import (
	"context"
	"core/helper"
	"errors"
	"model"
	"time"

	"core/internal/svc"
	"core/internal/types"
	"core/define"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// MailCodeSendRegister 邮箱注册发送验证码
func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.MailCodeSendRegisterRequest) (resp *types.MailCodeSendRegisterResponse, err error) {
	// todo: add your logic here and delete this line
	// 该邮箱不存在
	var ubs []model.UserBasic 
	cnt :=l.svcCtx.DB.Where("email=?",req.Email).Find(&ubs).RowsAffected // 查询邮箱是否已注册
	if err!=nil{
		return nil,errors.New("查询邮箱错误")
	}

	// 邮箱已注册
	if cnt>0{
		err = errors.New("该邮箱已注册")
		return nil,err
	}

	// 生成验证码
	code := helper.RandCode()
	// 存储验证码
	l.svcCtx.RDB.Set(l.ctx,req.Email,code,time.Second*time.Duration(define.CodeExpire))

	err=helper.MailSendCode(req.Email,code) // 发送验证码
	if err!=nil{
		return nil,errors.New("发送验证码错误")
	}
	return 

}
