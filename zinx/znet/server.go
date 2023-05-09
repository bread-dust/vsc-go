package znet

import (
	"fmt"
	"net"
	"ziface"
	"settings"
)

// iServer的接口实现，定义一个服务器模块
type Server struct {
	// 服务器的名称
	Name string
	// 服务器的绑定的IP版本
	IPVersion string
	// 服务器监听的IP
	IP string
	// 服务器监听的端口
	Port int
	// 当前的Server的消息管理模块，用来绑定MsgID和对应的处理业务API关系
	MsgHandle ziface.IMsgHandle


	// 该Server的链接管理器
	ConnMgr ziface.IConnManager

	// 该Server创建链接之后自动调用Hook函数--OnConnStart
	OnConnStart func(conn ziface.IConnection)
	// 该Server销毁链接之前自动调用Hook函数--OnConnStop
	OnConnStop func(conn ziface.IConnection)
}

// 路由功能
func (s *Server) AddRouter(msgId uint32,router ziface.IRouter) {
	s.MsgHandle.AddRouter(msgId, router)
	fmt.Println("Add Router succ!")
}

// 初始化Server模块的方法
func NewServer(name string) ziface.IServer {
	s := &Server {
		Name: name,
		IPVersion: settings.GlobalObject.IPVersion,
		IP:settings.GlobalObject.Host,
		Port: settings.GlobalObject.TcpPort, 
		MsgHandle: NewMsgHandle(),
		ConnMgr: NewConnManager(),
	}
	return s	
}

// 启动服务器
func (s *Server) Start() {
	fmt.Printf("[Zinx] server Name:%s, listenner at IP:%s, Port:%d is starting\n", settings.GlobalObject.Name, settings.GlobalObject.Host, settings.GlobalObject.TcpPort)
	
	// TODO
	go func(){
	
	// 0 开启消息队列及Worker工作池
	s.MsgHandle.StartWorkerPool()


	// 1. 获取一个Tcp的addr
	addr ,err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err!=nil{
		fmt.Println("resolve tcp addr error:", err)
		return
	}

	// 2. 监听服务器的地址
	listener, err:= net.ListenTCP(s.IPVersion, addr)
	if err!=nil{
		fmt.Println("listen", s.IPVersion, "err", err)
		return
	}

	fmt.Println("start Zinx server succ, ", s.Name, "succ, Listening...")

	var cid uint32
	cid = 0
	
	// 3. 阻塞客户端的连接请求，处理客户端的连接业务（读写）
	for{
		// 如果有客户端连接过来，阻塞会返回
		conn,err:=listener.AcceptTCP()
		if err!=nil{
			fmt.Println("Accept err", err)
			continue
		}
		
		// 判断当前的连接数是否超过最大值
		if s.ConnMgr.Len() >= settings.GlobalObject.MaxConn {
			conn.Close()
			continue
		}
		
		// 将处理新链接的业务方法和conn进行绑定，得到我们的链接模块
		deal := NewConnection(conn,cid, s.MsgHandle)
		cid ++

		// 启动当前的链接业务处理
		go deal.Start()
		}	
	}()
}
// 停止服务器
func (s *Server) Stop() {
	// TODO 将服务器的资源，状态或者一些已经开辟的链接信息 进行停止或者回收
	fmt.Println("[STOP] Zinx server name", s.Name)
	s.ConnMgr.ClearConn()
	

}

// 运行服务器
func (s *Server) Serve() {		
	// 启动server 的服务功能
	s.Start()
	// TODO 做一些启动服务器之后的额外业务

	// 阻塞状态
	select{}
}

func (s *Server) GetConnMgr()*ziface.IConnManager{
	return &s.ConnMgr
}

// 注册OnConnStart钩子函数的方法
func (s *Server) SetOnConnStart(hookFunc func(ziface.IConnection)) {
	s.OnConnStart = hookFunc
}

// 注册OnConnStop钩子函数的方法
func (s *Server) SetOnConnStop(hookFunc func(ziface.IConnection)) {
	s.OnConnStop = hookFunc
}

// 调用OnConnStart钩子函数的方法
func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("---> Call OnConnStart()...")
		s.OnConnStart(conn)
	}
}

// 调用OnConnStop钩子函数的方法
func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("---> Call OnConnStop()...")
		s.OnConnStop(conn)
	}
}
