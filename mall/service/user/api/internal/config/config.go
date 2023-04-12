package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql struct{ // 数据库配置
		DataSource string  // 连接地址 $user:$pwd@tcp(ip:port)
	}

	CacheRedis cache.CacheConf
 	Auth struct { // jwt鉴权配置
        AccessSecret string // jwt密钥
        AccessExpire int64 // 有效期，单位：秒
    }

}
