package main

import (
	"addserv1/pb"
	"context"
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

// gRPC test

// 使用bufconn 构建测试连接
const bufSize = 1024 * 1023
var bufListener *bufconn.Listener

func init(){
	bufListener = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	gs := NewGRPCServer(addService{})
	pb.RegisterAddServer(s,gs)
	go func ()  {
		if err:=s.Serve(bufListener);err!=nil{
			log.Fatal("server exit with error")
		}
	}()
}

func bufDialer(context.Context,string)(net.Conn,error){
	return bufListener.Dial()
}
// 可编写一个gRPC客户端，测试我们的gRPC是否正常

func TestSum(t *testing.T){
	// 建立连接
	conn,err:=grpc.DialContext(
		context.Background(),
		"bufnet",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(bufDialer),
	)
	if err!=nil{
		t.Fail()
	}
	defer conn.Close()

	c:=pb.NewAddClient(conn)
	resp,err := c.Sum(context.Background(),&pb.SumRequest{A:10,B:3})
	assert.Nil(t,err)
	assert.NotNil(t,resp)
	assert.Equal(t,resp.V,12)
}