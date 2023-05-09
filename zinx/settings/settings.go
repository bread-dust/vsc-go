package settings

import (
	"encoding/json"
	"io/ioutil"
	"ziface"


)

/*
	存储所有的全局变量
*/

type Config struct{
	// Server
	TcpServer ziface.IServer // 当前Zinx全局的Server对象
	Host string  // 当前服务器主机监听的IP
	TcpPort int // 当前服务器主机监听的端口号
	Name string // 当前服务器的名称
	IPVersion string // 当前服务器主机监听的IP版本

	// Zinx
	Version string// 当前Zinx的版本号
	MaxConn int // 当前服务器主机允许的最大连接数
	MaxPackageSize uint32// 当前Zinx框架数据包的最大值
	WorkerPoolSize uint32 // 当前业务工作Worker池的Goroutine数量
	MaxWorkerTaskLen uint32 // Zinx框架允许用户最多开辟多少个Worker(限定条件)
}

var GlobalObject=new(Config) // 定义一个全局的对象

func (g *Config) Reload(){
	// data,err:=ioutil.ReadFile("F:\\vsc_go\\zinx\\settings\\zinx.json")
	data,err:=ioutil.ReadFile("../../settings/zinx.json")

	if err!=nil{
		panic(err)
	}

	// 将json数据解析到struct中
	err = json.Unmarshal(data, &GlobalObject)
	if err!=nil{
		panic(err)
	}

}


// 初始化当前的GlobalObject
func init(){
	GlobalObject = &Config{
		Name: "ZinxServerApp",
		Version: "V0.4",
		TcpPort: 8999,
		Host: "127.0.0.1",
		MaxConn: 12000,
		MaxPackageSize: 4096,
		WorkerPoolSize: 10,
		MaxWorkerTaskLen: 1024,
	}

	GlobalObject.Reload()
}