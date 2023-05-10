package geerpc

import (
	"codec"
	"net"
)

const MagicNumber = 0x3bef5c

// 编码方式
type Option struct{
	MagicNumber int 
	CodecType codec.Type
}

var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodecType: codec.GobType,
}

type Server struct{}

func NewServer() *Server{
	return &Server{}
}

var DefaultServer = NewServer()

func (server *Server)Accept(lis net.Listener){
	for {
		conn,err := lis.Accept()
		if err != nil {
			log.Println("rpc server: accept error:",err)
			return
		}
		go server.ServerConn(conn)
	}
}

func Accept(lis net.Listener){
	DefaultServer.Accept(lis)
}