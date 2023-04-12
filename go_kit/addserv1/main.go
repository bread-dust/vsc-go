package main

import (
	"addserv1/pb"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"golang.org/x/sync/errgroup"

	"google.golang.org/grpc"
)

func main() {

	// 前置资源初始化
	srv := NewService()

	var g errgroup.Group
	// HTTP
	g.Go(func() error {
		httpListerner, err := net.Listen("tcp", "127.0.0.1:8081")
		if err!=nil{
			return err
		}
		defer httpListerner.Close()
		// 初始化logger	
		logger := log.NewLogfmtLogger(os.Stderr)
		httpHandler := NewHTTPServer(srv,logger)
		return http.Serve(httpListerner,httpHandler)
	})
	// gRPC
	g.Go(func() error {
		gs := NewGRPCServer(srv) // gRPC 服务模块
		listener, err := net.Listen("tcp", ":8082")
		if err != nil {
			fmt.Printf("err:%v\n", err)
			return err
		}
		s := grpc.NewServer() 
		pb.RegisterAddServer(s, gs)  // 一个处理器，两个服务数据穿透

		// 启动服务
		return s.Serve(listener)
	})

	if err:= g.Wait();err!=nil{
		fmt.Print("err")
	}
}
