package main

import (
	"dengliwei.com/add_client/proto"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 连接rpc server
	conn,err:=grpc.Dial("127.0.0.1:8973",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err !=nil {
		log.Fatalf("grpc.Dial failed,err:%v",err)
		return
	}
	defer conn.Close()
	
	// 创建rpc client
	client:=proto.NewCalcServiceClient(conn)

	// 发起rpc调用
	ctx,cancel:= context.WithTimeout(context.Background(),time.Second*2)
	defer cancel()
	retp,err:=client.Add(ctx,&proto.AddRequest{X: 10, Y: 20})
	if err!=nil{
		log.Fatal("client add failed ")
		return
	}
	log.Printf("ret:%v\n",retp.GetResult())

	

}