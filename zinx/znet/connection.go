package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"ziface"
)

/*
	链接模块
*/

type Connection struct{
	// 当前链接的socket TCP套接字
	Conn *net.TCPConn
	// 链接的ID
	ConnID uint32
	// 当前的链接状态
	isClosed bool
	// 告知当前链接已经退出/停止的channel
		ExitBuffChan chan bool

	// 消息管理MsgID和对应处理方法的消息管理模块
	MsgHandler ziface.IMsgHandle


}

// 初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, msgHanlder ziface.IMsgHandle) *Connection{
	c := &Connection{
		Conn: conn,
		ConnID: connID,
		MsgHandler: msgHanlder,
		isClosed: false,
		ExitBuffChan: make(chan bool, 1),
	}
	return c
}	

// 启动链接，让当前的链接准备开始工作
func (c *Connection)Start(){
	fmt.Println("Conn Start()...ConnID=", c.ConnID)
	// 启动从当前链接的读数据的业务
	go c.StartReader()
	// TODO 启动从当前链接写数据的业务
}

// 这里主要用于处理读数据的业务
func (c *Connection)StartReader(){
	fmt.Println("Reader Goroutine is running...")
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


		// 从路由中，找到注册绑定的Conn对应的router调用
		go c.MsgHandler.DoMsgHandler(&req)
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
	// 关闭socket链接
	c.Conn.Close()
	// 回收资源

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

// 提供一个SendMsg方法，将我们要发送给客户端的数据，先进行封包，再发送
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
	if _,err:=c.Conn.Write(binaryMsg);err!=nil{
		fmt.Println("Write msg id ",msgId," error")
		c.Stop()
		return errors.New("Conn Write error")
	}

	return nil

}
