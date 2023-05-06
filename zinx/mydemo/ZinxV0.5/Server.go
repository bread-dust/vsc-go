package main

import (
	"fmt"
	"ziface"
	"znet"
)

/*
	基于Zinx框架来开发的服务器端应用程序
*/

// ping test 自定义路由

type PingRouter struct {
	znet.BaseRouter
}

// Test PreHandle
func(this *PingRouter)PreHandle(request ziface.IRequest){
	fmt.Println("Call Router PreHandle...")

	// 先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId = ", request.GetMsgId()," data = ", string(request.GetData()))

	err := request.GetConnection().SendMsg(1,[]byte("ping...ping...ping"))
	if err!=nil{
		fmt.Println(err)
	}

}

// Test Handle
func(this *PingRouter)Handle(request ziface.IRequest){
	fmt.Println("Call Router HandleHandle...")
	_,err := request.GetConnection().GetTCPConnection().Write([]byte("ping... ping... \n"))
	
	if err != nil {
		fmt.Println("call back ping ping error")
	}
	
}
// Test PostHandle
func(this *PingRouter)PostHandle(request ziface.IRequest){
	fmt.Println("Call Router PreHandle...")
	_,err := request.GetConnection().GetTCPConnection().Write([]byte("after ping... \n"))
	
	if err != nil {
		fmt.Println("call back after ping error")
	}
	
}

func main() {
	// 1. 创建一个server句柄，使用Zinx的api
	s := znet.NewServer("Zinx V0.4")
	// 2. 给当前zinx框架添加一个自定义的router
	s.AddRouter(&PingRouter{})
	// 3. 启动server
	s.Serve()
}