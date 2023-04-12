package main

import (
	"addserv1/pb"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
)

// transport

// 网络传输相关，包括协议（HTTP，gRPC，thrift...）等
// 3. transport
// decode
// 请求来了之后 根据协议和编码 去解析数据
// http json
func decodeSumRequest(ctx context.Context,r *http.Request)(interface{},error){
	var request SumRequest
	if err:=json.NewDecoder(r.Body).Decode(&request);err!=nil{ // 将请求解析数据后，存放到SumRequest 中
		return nil,err
	}
	return request,nil
}

func decodeConcatRequest(ctx context.Context,r *http.Request)(interface{},error){
	var request ConcatRequest
	if err:=json.NewDecoder(r.Body).Decode(&request);err!=nil{ 
		return nil,err
	}
	return request,nil
}

// 编码
// 把响应数据 按协议和编码 返回
// w是响应的网络句柄，response 是业务层响应的数据
func encodeResponse(ctx context.Context,w http.ResponseWriter,response interface{}) error {
	return json.NewEncoder(w).Encode(response) 
}

func NewHTTPServer(srv Addservice,logger log.Logger)http.Handler{ // 对sum服务的http请求处理器

	sum := makeSumEndponit(srv)
	// log。With() 派生子日志
	sum = loggingMiddleware(log.With(logger,"method","sum"))(sum)
	// 限流中间件
	sum = reateMiddleware(rate.NewLimiter(1,1))(sum)
	sumHandler := httptransport.NewServer(
		sum,  // 日志中间件包一层的endpoint
		decodeSumRequest,
		encodeResponse,
	)

	concatHandler := httptransport.NewServer( // 对concat 服务的 http请求处理器
		makeConcatEndponit(srv),
		decodeConcatRequest,
		encodeResponse,
	)

	http.Handle("/sum",sumHandler)
	http.Handle("/concat",concatHandler)
	http.ListenAndServe(":8081",nil)
	
	// gin
	r := gin.Default()
	r.POST("/suM",gin.WrapH(sumHandler))
	r.POST("/concat",gin.WrapH(concatHandler))

	return r

}

// gRPC
// 请求与响应// NewGRPCServer 构造函数
// grpc
// grpcServer 分为两部分 外侧gRPC 定义对外进出进出的请求和响应+内侧go-kit 
type grpcServer struct{ // gRPC 服务模块
	pb.UnimplementedAddServer

	// 有两个gRRC处理器
	sum grpctransport.Handler
	concat grpctransport.Handler
}


func NewGRPCServer(svc Addservice) pb.AddServer{
	return &grpcServer{
		sum: grpctransport.NewServer( // 对sum 服务的gRPC 请求处理器
			makeSumEndponit(svc), //endpoint
			decodeGRPCSumRequest,
			encodeGRPCSumResponse,
		),
		concat: grpctransport.NewServer(
			makeConcatEndponit(svc),
			decodeGRPCConcatRequest,
			encodeGRPCConcatResponse,
		),
	}
}

func (s grpcServer) Sum(ctx context.Context,req *pb.SumRequest) (*pb.SumResponse,error){ //接口实现
	_,resp,err := s.sum.ServeGRPC(ctx,req) 
	if err!=nil{
		return nil,err
	}
	return resp.(*pb.SumResponse),nil
}

func (s grpcServer) Concat(ctx context.Context,req *pb.ConcatRequest) (*pb.ConcatResponse,error){  //接口实现
	_,resp,err := s.concat.ServeGRPC(ctx,req) // 交给go—kit处理
	if err!=nil{
		return nil,err
	}
	return resp.(*pb.ConcatResponse),nil
}

func decodeGRPCSumRequest(_ context.Context,grpcReq interface{})(interface{},error){
	req := grpcReq.(*pb.SumRequest)
	return SumRequest{
		A: int(req.A),
		B: int(req.B),
	},nil
}

func decodeGRPCConcatRequest(_ context.Context,grpcReq interface{})(interface{},error){
	req :=grpcReq.(*pb.ConcatRequest)
	return ConcatRequest{
		A: req.A,
		B: req.B,
	},nil
}

func encodeGRPCSumResponse(_ context.Context,response interface{})(interface{},error){
	resq := response.(*pb.SumResponse)
	return pb.SumResponse{
		V:   int64(resq.V),
		Err: resq.Err,
	},nil
}

func encodeGRPCConcatResponse(_ context.Context,response interface{})(interface{},error){
	resp :=response.(*pb.ConcatResponse)
	return pb.ConcatResponse{
		V:   resp.V,
		Err: resp.Err,
	},nil
}
