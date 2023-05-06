package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"settings"
	"ziface"
)

// 封包、拆包模块
type DataPack struct{}

// 拆包封包实例的一个初始化方法
func NewDataPack() *DataPack {
	return &DataPack{}
}

// 获取包的头的长度方法
func (dp *DataPack)GetHeadLen() uint32{
	// Id uint32(4字节) + DataLen uint32(4字节)
	return 8
}
// 封包方法
func (dp *DataPack)Pack(msg ziface.IMessage)([]byte, error){
	// 创建一个存放bytes 字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	// 将dataLen写进dataBuff中
	if err:= binary.Write(dataBuff,binary.LittleEndian,msg.GetMsgLen());err!=nil{
		return nil,err
	}
	// 将MsgId写进dataBuff中
	if err:= binary.Write(dataBuff,binary.LittleEndian,msg.GetMsgId());err!=nil{
		return nil,err
	}

	// 将data数据写进dataBuff中
	if err:= binary.Write(dataBuff,binary.LittleEndian,msg.GetData());err!=nil{
		return nil,err
	}
	return dataBuff.Bytes(), nil
}

// 拆包方法
func (dp *DataPack)Unpack(binaryData []byte)(ziface.IMessage, error){
	// 创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)

	// 只解压head信息，得到dataLen和msgId
	msg := &Message{}

	// 读dataLen
	if err:=binary.Read(dataBuff,binary.LittleEndian,&msg.DataLen);err!=nil{
		return nil,err
	}
	// 读msgId
	if err:=binary.Read(dataBuff,binary.LittleEndian,&msg.Id);err!=nil{
			return nil,err
	}

	// 判断dataLen的长度是否超出我们允许的最大包长度
	if (settings.GlobalObject.MaxPackageSize>0&&msg.DataLen>settings.GlobalObject.MaxPackageSize){
		println(msg.DataLen)
		return nil,errors.New("too large msg data recv!")
	}

	return msg,nil
}
