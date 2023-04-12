package config

import "github.com/zeromicro/go-zero/rest"

// Config is the configuration.
type Config struct {
	rest.RestConf // 继承rest配置

	// Mysql is the mysql configuration.
	Mysql struct{
		DataSource string 
	}
	
	// Redis is the redis configuration.
	Redis struct{
		Addr string
		Pass string
	}
}
