package main

import (
	"add_server/proto"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{
	proto.UnimplementedCalcServiceServer
}

func (s server)Add(ctx context.Context,in *proto.AddRequest) (*proto.AddReponse, error) {
	sum := int64(in.GetX())+int64(in.GetY())
	return &proto.AddReponse{Result:sum},nil
}
	
func main() {
	l,err:=net.Listen("tcp",":8973")
	if err!=nil{
		log.Fatalf("net.listen failed,err:%v",err)
		return
	}
	s:=grpc.NewServer()
	proto.RegisterCalcServiceServer(s,&server{})
	//注册
	s.Serve(l)
	// 启动服务

}