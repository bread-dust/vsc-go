package main

import (
	"bookstore/pb"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	GRPCLISTEN = ":8972"
	HTTPLISTEN = ":8090"
)
func main() {
	// 连接数据库
	db,err := NewDB()
	if err != nil {
		fmt.Println("connect mysql failedo")
		return
	}

	// 创建server
	srv := server{
		bs:&bookstore{db:db},
	}
	//定义gRPC服务
	//监听端口
	l,err:=net.Listen("tcp",GRPCLISTEN)
	if err!=nil{
		fmt.Printf("failed listen:%v",err)
		return
	}
	// 创建一个grpc server 对象
	s:=grpc.NewServer()
	// 注册服务
	pb.RegisterBookstoreServer(s,&srv)
	
	// gRPC-Gateway gwmux:grpc的多路复用路由
	// 新建一个grpc handler gwmux将请求交给GRPCLISTEN处理
	gwmux := runtime.NewServeMux()
	dops:= []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = pb.RegisterBookstoreHandlerFromEndpoint(context.Background(),gwmux,GRPCLISTEN,dops)
	if err!=nil{
		log.Fatal("failed to register service to gwmux")
	}
	//新建一个http handler复用路由，将http请求交给grpc的复用路由处理，mux:http->gprc
	mux := http.NewServeMux()
	mux.Handle("/",gwmux)

	// 定义HTTP server 配置
	gwServer := &http.Server{
		Addr: HTTPLISTEN,
		Handler: grpcHandlerFunc(s,mux), //请求的统一入口
	}
	// 开始grpc_gateway服务
	fmt.Println("grpc-gateway serve on 8090")
	gwServer.Serve(l)
}

// 将gRPC 请求和HTTP请求分别调用不同的handler处理
func grpcHandlerFunc(grpcServer *grpc.Server,otherHandler http.Handler) http.Handler{
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 检测协议 判断'content-type'是否为gprc
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-type"),"application/gprc"){
			grpcServer.ServeHTTP(w,r)
		}else {
			otherHandler.ServeHTTP(w,r)
		}
	}),&http2.Server{})
}