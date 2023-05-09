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

// 创建链接之后执行的钩子函数
func DoConnectionBegin(conn ziface.IConnection){
	fmt.Println("===> DoConnectionBegin is Called ...")
	if err := conn.SendMsg(202,[]byte("DoConnection BEGIN")); err != nil{
		fmt.Println(err)
	}

	// 给当前的连接设置一些属性
	fmt.Println("Set conn Name, Home done!")
	conn.SetProperty("Name","Zinx V0.9")
	conn.SetProperty("Home","")
	conn.SetProperty("Time","2020-04-18")
}

// 链接销毁之前执行的钩子函数
func DoConnectionLost(conn ziface.IConnection){
	fmt.Println("===> DoConnectionLost is Called ...")
	fmt.Println("conn ID = ",conn.GetConnID()," is Lost ...")
	// 获取链接属性
	if name,err := conn.GetProperty("Name");err == nil{
		fmt.Println("Name = ",name)
	}
	if home,err := conn.GetProperty("Home");err == nil{
		fmt.Println("Home = ",home)
	}
	if time,err := conn.GetProperty("Time");err == nil{
		fmt.Println("Time = ",time)
	}
}


func main() {
	// 1. 创建一个server句柄，使用Zinx的api
	s := znet.NewServer("Zinx V0.4")
	// 2. 给当前zinx框架添加一个自定义的router
	s.AddRouter(0,&PingRouter{})
	// 3. 启动server
	s.Serve()
}