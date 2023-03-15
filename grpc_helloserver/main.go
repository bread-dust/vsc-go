package main

import (
	"fmt"
	"grpc_helloserver/pb"
	"io"
	"net"

	"google.golang.org/grpc"
)

type server struct{
	pb.UnimplementedGreeterServer
}

func (s *server) LotsOfReplies(in *pb.HelloRequest,stream pb.Greeter_LotsOfRepliesServer) error {
	words := []string{
		"你好",
		"hello",
		"こんにちは",
		"안녕하세요",
	}
	for _,word := range words{
		data := &pb.HelloResponse{
			Reply: word + in.GetName(),
		}
		// 使用Send 返回多个数据，发送一个stream resp
		if err:=stream.Send(data);err!=nil{
			return err
		}
	}
	return nil
}

func (s *server) LotsOfGreeting(stream pb.Greeter_LotsOfGreetingServer) error{
	reply := "你好:"
	for{
		//接收客户端发来的流数据（将队列的req读取）
		res,err:=stream.Recv()
		if err==io.EOF{
			// 发送resp流
			return stream.SendAndClose(&pb.HelloResponse{
				Reply: reply,
			})
		}
		if err!=nil{
			return err
		}
		reply += res.GetName()
	}
}

func main() {
		// 监听本地的8972端口
		lis, err := net.Listen("tcp", ":8972")
		if err != nil {
			fmt.Printf("failed to listen: %v", err)
			return
		}
		s := grpc.NewServer()                  // 创建gRPC服务器
		pb.RegisterGreeterServer(s, &server{}) // 在gRPC服务端注册服务
		// 启动服务
		err = s.Serve(lis)
		if err != nil {
			fmt.Printf("failed to serve: %v", err)
			return
		}
}