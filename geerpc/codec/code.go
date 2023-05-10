package codec

import "io"

// rpc Header 请求头
type Header struct{
	ServiceMethod string //format:"Service.Method" 服务名和方法名
	Seq uint64 //sequence number by client 请求序号
	Error error
}

// Codec 编解码器
type Codec interface{
	io.Closer
	ReadHeader(*Header) error //读取请求头
	ReadBody(interface{}) error //读取请求体
	Write(*Header,interface{}) error //写入请求头和请求体
}

type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

// 两种编码类型
const (
	GobType Type = "application/gob" 
	JsonType Type = "application/json" 
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init()  {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}