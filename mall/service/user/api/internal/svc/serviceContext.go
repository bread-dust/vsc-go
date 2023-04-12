package svc

import (
	"api/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"

	"model"
)

type ServiceContext struct {
	Config config.Config

	UserModel model.UserModel //加入User表增删改查操作Model
	Cost rest.Middleware //自定义中间件
}

func NewServiceContext(c config.Config) *ServiceContext {
	// UserModel  -> 接口类型
	// *defaultUserModel 实现了接口
	// NewUserModel 
	//需要传一个sql连接
	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		UserModel:model.NewUserModel(sqlxConn,c.CacheRedis),
		Cost:middleware.NewCostMiddleware().Handle,
	}
}
