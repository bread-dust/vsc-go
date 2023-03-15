package main

import (
	"context"
	"flag"
	"grpc_helloclient/pb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var name = flag.String("name", "dengliwei", "name to greet")
var addr = flag.String("addr","127.0.0.1:8972","the address to connect to")

func runLotsOfReplies(c pb.GreeterClient)  {
	//server 端流式rpc
	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*2)
	defer cancel()
	// 调用方法，传一个值req给服务端，返回一个strem resp(创建一个resp队列)
	stream,err:=c.LotsOfReplies(ctx,&pb.HelloRequest{Name:*name})
	if err!=nil{
		log.Fatalf("LotsOfReplies failed,err:%v\n",err)
	}
	for {
		// 接收服务端返回的流式数据（从队列中拿出来resp）
		res,err:=stream.Recv()
		if err==io.EOF{
			break
		}
		if err!=nil{
			log.Fatalf("got reply:%q\n",res.GetReply())
		}
		log.Printf("got reply:%q\n",res.GetReply())
	}

}

func runLotsOfGreetings(c pb.GreeterClient)  {
	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*2)
	defer cancel()
	stream ,err := c.LotsOfGreetings(ctx)
	if err!=nil{
		log.Fatal("c.LotsOfGreetings failed")
	}
	names:=[]string{"dengliwei","zhangsan"}
	for _,name := range names{
		err := stream.Send(&pb.HelloRequest{Name: name})
		if err!=nil{log.Fatal("Streamsend failed")
			return
		}
	}
	//关闭发送流，将resp流接收
	res,err:= stream.CloseAndRecv()
	if err!=nil{
		log.Fatal("greeting failed ")
	}
	log.Printf("got reply:%v",res.GetReply())
}
	

// grpc 客户端
// 调用server 端的sayHello方法
func main() {
	flag.Parse()
	// 连接server 端
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc dial failed,err:%v", err)
		return
	}
	defer conn.Close()

	// 创建客户端
	c := pb.NewGreeterClient(conn) 
	// 调用流式rpc方法
	runLotsOfReplies(c)

	runLotsOfGreetings(c)
}
