package codec

import "io"

type Header struct{
	ServiceMethod string //format:"Service.Method" 服务名和方法名
	Seq uint64 //sequence number by client 请求序号
	Error error
}

type Codec interface{
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header,interface{}) error
}

type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

const (
	GobType Type = "application/gob" //定义两种Codec	
	JsonType Type = "application/json"
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init()  {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}