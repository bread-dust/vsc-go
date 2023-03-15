package main

import (
	"encoding/base64"
	"encoding/json"
)

// 基于游标的分页

type Page struct{
	NextID string `json:"next_id"`
	NextTimeAtUTC int64 `json:"next_time_at_utc"`
	PageSize int64 `json:"page_size"`
}

type Token string

// Encode 返回分页token
func (p Page) Encode() Token{
	// 序列化成字节
	b,err := json.Marshal(p)
	if err!=nil{
		return Token("")
	}
	//字节转成bsase64编码字符串
	return Token(base64.RawStdEncoding.EncodeToString(b))
}

// Decode 解析分页信息
func (t Token) Decode() Page{
	var result Page
	if len(t) == 0{
		return result
	}

	// base64 string-> byte
	bytes,err:=base64.StdEncoding.DecodeString(string(t))
	if err!=nil{
		return result
	}
	// 反序列化成json,bype->json
	err = json.Unmarshal(bytes,&result)
	if err!=nil{
		return result
	}
	return result 
}