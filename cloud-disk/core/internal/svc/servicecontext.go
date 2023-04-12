package svc

import (
	"core/internal/config"
	"core/internal/middleware"
	"model"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

// ServiceContext is the service context.
type ServiceContext struct {
	Config config.Config // 配置文件
	DB *gorm.DB // 数据库
	RDB *redis.Client // redis
	Auth rest.Middleware // 鉴权中间件

}

// NewServiceContext creates a new ServiceContext.
func NewServiceContext(c config.Config) *ServiceContext {
	// 
	return &ServiceContext{ // 返回服务上下文
		Config: c, // 配置文件
		DB: model.InitMysql(c.Mysql.DataSource), // 初始化数据库
		RDB: model.InitRedis(c.Redis.Addr,c.Redis.Pass), // 初始化redis
		Auth: middleware.NewAuthMiddleware().Handle, // 初始化鉴权中间件
		
	}
}
