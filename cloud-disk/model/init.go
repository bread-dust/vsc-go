package model

import (
	"log"

	"gorm.io/driver/mysql"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)


// 初始化数据库
func InitMysql(DataSource string)*gorm.DB{

	// gorm创建数据库引擎 
	db,err:=gorm.Open(mysql.Open(DataSource),&gorm.Config{})
	// 判断是否出错
	if err!=nil{
		log.Printf("Gorm new engine:%v\n",err)
		return nil
	}
	return db
}

// 初始化redis
func InitRedis(Addr string,Pass string)*redis.Client{
	// 创建redis客户端
	var rdb =redis.NewClient(&redis.Options{
	Addr:Addr	,
	Password: Pass,
	DB: 0, // use default DB
})
	return rdb

}