package main

import (
	"context"
	"errors"
	"io"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/log"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/genproto/googleapis/cloud/retail/v2"
	"google.golang.org/grpc"
)

// Endpoint中间件

func loggingMiddleware(logger log.Logger)endpoint.Middleware{
	return func(next endpoint.Endpoint)endpoint.Endpoint{
		return func(ctx context.Context,request interface{})(interface{},error){
			logger.Log("msg","calling endpoint")
			defer logger.Log("msg","called endpint")
			return next(ctx,request)
		}
	}
}

// 限流中间件
// golang.rog/x/time/rate
func rateMidddleware(limit *rate.Limit)endpoint.Middleware{
	return func(next endpoint.Endpoint)endpoint.Endpoint{
		return func(ctx context.Context,request interface{})(interface{},error){
			// 限流逻辑
			if limit.Allow(){
				return next(ctx,request)
			}else{
				return nil,errors.New("error")
			}
		}
	}	

}


// Service中间件
type ServiceMiddleware func(Service) Service //定义中间件处理器

type loggingMiddlewareService struct{  // 定义中间件服务
	logger *zap.Logger 
	next Service
}

func NewService(log *zap.Logger)*ServiceMiddleware{
	return func (next Service)Service  {
		return &loggingMiddlewareService{
			logger:log,
			next next
		}
	}
}
//新的 loggingmiddlewareService实现所有方法


// go-kit client 客户端



type trimRequest struct {
	string a=1;
}

type trimResponse struct {
	 string b=1;
}


// -- endpoint层--
// 不直接调用serivce层，提供服务，而是请求其他服务
// grpctansport "github.com/go-kit/kit/transpot/grpc"
func makeTrimEndpoint(conn *grpc.Clientconn)endpoint{
	return grpctransport.NewClient（
		conn, //grpc conn
		"pb.Trim", // service name
		"TrimSpace", // method
)
}

//- transport层--
func encodeTrimRequest(_ context.Context,req interface{})(interface{},error)  {
	req := req.(trimRequest)
	return &pb.trimRequset{
		a:req.a
	},nil
}

func decodeTrimReponse(_ context.Context,resp interface{})(interface{},error)  {
	resp := resp.(pb.TrimResponse)
	return &trimResponse{b:resp.b},nil
}

// -- service 层
type Client struct{
	old Service
	trim endpoint.Endpoint
}

func NewClient(trimpoint endpoint.Endpoint,oldservice Service)Service{
	return &Client{
		old:oldservice,
		trim:trimpoint
	}
}

// 实现新的方法
func (c Client)Sum(ctx,a,b)  {
	return c.old.Sum(ctx,a,b)
}

func (c Client)Concat(){
	//先调用外部trim-service 对数据进行处理
	resp_A , err:= c.trim(ctx,trimRequest{a:ad}) //作为客户端对外请求
	if err!-nil=nil{
		return "",nil
	}
	resp_B , err:= c.trim(ctx,trimRequesr{a:aa}) //作为客户端对外请求
	if err!-nil=nil{
		return "",nil
	}
	trimA:=resq_A.(trimResponse)
	trimB:+resq_B.(trimResponse)

	return trim.old.Sum(ctx,trimA.s,trimB.s)
}
// -- main.go --
conn,err := grpc.Dial()
defer conn.Close()
trimEndpoint := makeTrimEndpoint(conn)
src := NewClient(trimEndpoint,oldsrc)	

//consul 
// 从注册中心获取服务的地址
// 基于consul 实现对trim service的实现
// consulapi "github.com/hashicorp/consul/api"
// sdconsul "github.com/go-kit/kit/consul"


func getTrimServiceFromConsul(consulAddr string,logger log.logger,srvName string,tags []string) (endpoint.Endpoint.error) {
	// 连接consul
	cfg := consulapi.DefaultConfig()
	cfg.Address = consulAdddr
	consulClient,err:=consulapi.NewClient(cfg)

	// 使用go-kit适配器包装
	sdClinet := sdconsul.NewClient(consulClient)
	instancer := sdconsul.NewInstancer(sdClient,logger,srvName,tags,true) // 日志，服务名称，过滤标签

	// endpointer
	endpointer := sd.NewEndpointer(instancer,factory,logger)

	// banlancer
	balancer := lb.NewRoundRobin(endpointer)

	// retry
	retry := lb.Retry(3,time.Second,balancer)
	return retry,nil
	
}

func Factory(instancer string) (endpoint.Endpoint,io.Closer,error) {
	conn,err:= grpc.Dial(instancer,grpc.WithInsecure())
	e:= makeTrimEndpoint(conn)
	retrun e,conn,err
}
// --main--
trimpointer,err := getTrimServiceFromConsul("localhost:850","trim_service",nil,logger,)
srv = NewClient(trimpointer,srv)