package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"settings"
	"sync"
	"ziface"
)

/*
	链接模块
*/

type Connection struct{

	// 当前Conn隶属于哪个Server
	TcpServer ziface.IServer

	// 当前链接的socket TCP套接字
	Conn *net.TCPConn
	// 链接的ID
	ConnID uint32
	// 当前的链接状态
	isClosed bool
	// 告知当前链接已经退出/停止的channel
	ExitBuffChan chan bool

	// 无缓冲管道，用于读、写Goroutine之间的消息通信
	msgChan chan []byte

	// 消息管理MsgID和对应处理方法的消息管理模块
	MsgHandler ziface.IMsgHandle

	// 链接属性集合
	property map[string]interface{}

	// 保护链接属性的锁
	propertyLock sync.RWMutex

}

// 初始化链接模块的方法
func NewConnection(server ziface.IServer,conn *net.TCPConn, connID uint32, msgHanlder ziface.IMsgHandle) *Connection{
	c := &Connection{
		TcpServer: server,
		Conn: conn,
		ConnID: connID,
		MsgHandler: msgHanlder,
		isClosed: false,
		ExitBuffChan: make(chan bool, 1),
		msgChan: make(chan []byte),
	}
	// 将conn加入到ConnManager中
	c.TcpServer.GetConnMgr().Add(c)

	return c
}	

// 写消息Goroutine， 用户将数据发送给客户端
func (c *Connection) StartWriter(){
	fmt.Println("[Writer Goroutine is running]")
	defer fmt.Println(c.RemoteAddr().String(), "[conn Writer exit!]")
	// 不断的阻塞的等待channel的消息，进行写给客户端
	for{
		select {
		case data:= <- c.msgChan:
			// 有数据要写给客户端
			if _,err:=c.Conn.Write(data);err!=nil{
				fmt.Println("Send data error:, ",err," Conn Writer exit")
				return
			}
		case <- c.ExitBuffChan:
			// Reader已经退出，此时Writer也要退出
			return

		}
	}

}


// 启动链接，让当前的链接准备开始工作
func (c *Connection)Start(){
	fmt.Println("Conn Start()...ConnID=", c.ConnID)
	// 启动从当前链接的读数据的业务
	go c.StartReader()
	// TODO 启动从当前链接写数据的业务
	go c.StartWriter()

	// 按照开发者传递进来的 创建链接之后需要调用的处理业务，执行对应的Hook函数
	c.TcpServer.CallOnConnStart(c)
}

// 这里主要用于处理读数据的业务
func (c *Connection)StartReader(){
	fmt.Println("[Reader Goroutine is running...]")
	defer fmt.Println("connID = ", c.ConnID, " Reader is exit, remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		// 读取客户端的数据到buf中，最大512字节
		// buf := make([]byte, settings.GlobalObject.MaxPackageSize)
		// _, err := c.Conn.Read(buf)
		// if err != nil{
		// 	fmt.Println("recv buf err", err)
		// 	continue
		// }

		// 创建一个拆包解包对象
		dp := NewDataPack()

		// 读取客户端的Msg Head 二进制流 8个字节
		headData := make([]byte, dp.GetHeadLen())
		_,err:=io.ReadFull(c.GetTCPConnection(),headData)
		if err!=nil{
			fmt.Println("read head error")
			break
		}

		// 得到msgID 和 msgDatalen 放在msg消息中

		msg,err:=dp.Unpack(headData)
		if err!=nil{
			fmt.Println("unpack error",err)
			break
		}
		
		// 根据dalen 再次读取Data， 放在msg.Data中
		var data []byte
		if msg.GetMsgLen()>0{
			data=make([]byte,msg.GetMsgLen())
			_,err:=io.ReadFull(c.GetTCPConnection(),data)
			if err!=nil{
				fmt.Println("read msg data error",err)
				break
			}
		}

		msg.SetData(data)

		req := Request{
			conn: c,
			msg: msg,
		}


		if settings.GlobalObject.WorkerPoolSize>0{
			// 已经开启了工作池机制，将消息交给Worker处理
			c.MsgHandler.SendMsgToTaskQueue(&req)
		}else{
			// 根据绑定好的MsgID找到对应处理api业务执行
			go c.MsgHandler.DoMsgHandler(&req)
		}


	}
}


// 停止链接，结束当前链接的工作
func (c *Connection)Stop(){
	fmt.Println("Conn Stop()...ConnID=", c.ConnID)
	// 如果当前链接已经关闭
	if c.isClosed {
		return
	}
	c.isClosed = true

	// 调用开发者注册的 销毁链接之前 需要执行的业务Hook函数
	c.TcpServer.CallOnConnStop(c)


	// 关闭socket链接
	c.Conn.Close()
	// 回收资源
	close(c.ExitBuffChan)
	close(c.msgChan)

	// 告知Writer关闭
	c.ExitBuffChan <- true

	// 将当前链接从ConnMgr中摘除掉
	c.TcpServer.GetConnMgr().Remove(c)
}

// 获取当前链接的绑定socket conn
func (c *Connection)GetTCPConnection() *net.TCPConn{
	return c.Conn
}

// 获取当前链接模块的链接ID
func (c *Connection)GetConnID() uint32{
	return c.ConnID
}

// 获取远程客户端的TCP状态 IP port
func (c *Connection)RemoteAddr() net.Addr{
	return c.Conn.RemoteAddr()
}

// SendMsg 提供一个SendMsg方法，将我们要发送给客户端的数据，先进行封包，再发送
func (c *Connection) SendMsg(msgId uint32 ,data []byte)error{
	if c.isClosed{
		return errors.New("Connection closed when send msg")
	}
	
	// 将data进行封包 MsgDataLen|MsgID|Data
	dp:=NewDataPack()
	binaryMsg,err:=dp.Pack(NewMessage(msgId,data))
	if err!=nil{
		fmt.Println("Pack error msg id = ",msgId)
		return errors.New("Pack error msg")
	}

	// 将数据发送给客户端
	c.msgChan<-binaryMsg
	return nil

}

// 将消息交个TaskQueue，由worker进行处理
func (mh *MsgHandle) SendMsgToTaskQueue(reuquest ziface.IRequest){
	// 1. 将消息平均分配给不同的worker
	// 根据ConnID来进行分配
	workerID:=reuquest.GetConnection().GetConnID()%mh.WorkerPoolSize
	fmt.Println("Add ConnID=",reuquest.GetConnection().GetConnID()," request msgID=",reuquest.GetMsgId()," to workerID=",workerID)
	// 2. 将消息发送给对应的worker的TaskQueue即可
	mh.TaskQueue[workerID]<-reuquest
}

// 设置链接属性
func (c *Connection)SetProperty(key string,value interface{}){
	// 保护共享资源Map 加写锁
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	// 添加一个链接属性
	c.property[key]=value
	fmt.Println("Set Property key=",key," value=",value)
}
// 获取链接属性
func (c *Connection)GetProperty(key string)(interface{},error){
	// 保护共享资源Map 加读锁
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()

	// 读取属性
	if value,ok:=c.property[key];ok{
		return value,nil
	}else{
		return nil,errors.New("no property found")
	}
}

// 移除链接属性
func (c *Connection)RemoveProperty(key string){
	// 保护共享资源Map 加写锁
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	// 删除属性
	delete(c.property,key)
	fmt.Println("Remove Property key=",key)
}