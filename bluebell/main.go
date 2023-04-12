/*
@author:Deng.l.w
@version:1.20
@date:2023-02-23 18:09
@file:main.go
*/

package main

import (
	"controllers"
	"dao/mysql"
	"dao/redis"
	"logger"
	"pkg/snowflake"
	"routes"
	"settings"
	"context"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Go web 开发通用脚手架

func main() {
	// 1. 加载配置文件
	if len(os.Args) < 2 {
		fmt.Println("need a parameter")
		return
	}
	if err := settings.Init(os.Args[1]); err != nil {
		fmt.Printf("init settings err:%v\n", err)
	}

	// 2. 初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.AppConfig.Mode); err != nil {
		fmt.Printf("init logger err:%v\n", err)
		return
	}
	// 使用全局变量
	zap.L().Debug("logger init success")

	// 3. 初始化MySQL
	if err := mysql.Init(settings.Conf.MySOLConfig); err != nil {
		fmt.Printf("init mysql err:%v\n", err)
	}
	defer mysql.Close()
	// 4. 初始化 Redis

	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis err:%v\n", err)
	}
	defer redis.Close()

	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Printf("init snowflade failed:%v\n", err)
		return
	}
	// 初始化校验器使用的翻译器
	if err := controllers.InitTrans("zh"); err != nil {
		fmt.Printf("init validator failed,err:%v\n", err)
		return
	}
	// 5.注册路由
	r := routes.Setup()

	// 6. 启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.AppConfig.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
