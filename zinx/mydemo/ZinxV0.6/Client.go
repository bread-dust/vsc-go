package main

import (
	"fmt"
	"io"
	"net"
	"time"

	"znet"
)

// 模拟客户端
func main() {
	fmt.Println("client start...")

	time.Sleep(1*time.Second)
	// 直接连接远程服务器，得到一个conn连接
	conn,err:=net.Dial("tcp", "127.0.0.1:8999")
	if err!=nil{
		fmt.Println("client start err, exit!")
		return
	}
	// 连接调用Write方法，写数据
	for{
		
		// 发送封包的Msg消息
		dp := znet.NewDataPack()
		binaryMsg,err:=dp.Pack(znet.NewMessage(0,[]byte("ZinxV0.5 client Test Message")))

		if err!=nil{
			fmt.Println("Pack error msg id = ",0)
			return
		}

		if _,err:=conn.Write(binaryMsg);err!=nil{
			fmt.Println("write error err ",err)
					return
		}
		
		// 服务器就应该给我们回复一个message数据，MsgID:1 ping...ping...ping
		// 1 先读取流中的head部分，得到ID和dataLen
		binaryHead := make([]byte,dp.GetHeadLen())
		if _,err:=conn.Read(binaryHead);err!=nil{
			fmt.Println("read head error")
			break
		}
		// 将二进制的head拆包到msg结构体中
		msgHead,err:=dp.Unpack(binaryHead)
		if err!=nil{
			fmt.Println("server unpack err:",err)
			return
		}

		if msgHead.GetMsgLen()>0{
			msg:= msgHead.(*znet.Message)
			msg.Data = make([]byte,msg.GetMsgLen())

			if _,err:= io.ReadFull(conn,msg.Data);err!=nil{
				fmt.Println("server unpack data err:",err)
				return
			}

			fmt.Println("==> Recv Msg: ID=",msg.Id,"len=",msg.DataLen,"data=",string(msg.Data))
		}


		// cpu阻塞
		time.Sleep(1*time.Second)

	}
}