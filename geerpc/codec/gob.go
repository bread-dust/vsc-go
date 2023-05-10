package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
)


type GobCodec struct {
	conn io.ReadWriteCloser // 链接实例,存放tcp字节流
	buf *bufio.Writer //防止阻塞带缓冲的writer，可以提升性能
	dec *gob.Decoder //解码器
	enc *gob.Encoder //编码器
}


// 确保GobCodec实现了Codec接口
var _ Codec = (*GobCodec)(nil)

// NewGobCodec 创建一个编解码器
func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buf :=bufio.NewWriter(conn)
	return &GobCodec{
		conn: conn,
		buf:  buf,
		dec:  gob.NewDecoder(conn),
		enc:  gob.NewEncoder(buf),
	}
}

// ReadHeader 读取请求头
func (c *GobCodec) ReadHeader (h *Header) error  {
	return c.dec.Decode(h)
}

// ReadBody 读取请求体
func (c *GobCodec) ReadBody (body interface{}) error  {
	return c.dec.Decode(body)
}

// Write 写入请求头和请求体
func (c *GobCodec) Write (h *Header,body interface{}) (err error) {
	defer func(){
	_ = c.buf.Flush()
	if err!=nil{
		_ = c.conn.Close()
	}
	}()
	if err:= c.enc.Encode(h);err!=nil{
		log.Println("rpc codec:gob error encoding header:",err)
		return err
	}
	if err:= c.enc.Encode(body);err!=nil{
		log.Println("rpc codec : gob erorr encoding body:",err)
		return err
	}
	return nil
}

// Close 关闭链接
func (c *GobCodec) Close() error {
	return c.conn.Close()
}