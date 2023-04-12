package middleware

import (
	"bytes"
	"fmt"
	"net/http"
)

type bodyCopy struct{
	http.ResponseWriter // 结构体嵌套
	body *bytes.Buffer //记录响应内容
}

func NewBodyCopy(w http.ResponseWriter)*bodyCopy{
		return &bodyCopy{
			ResponseWriter: w,
			body:           bytes.NewBuffer([]byte{}),

		}
}
func (bc bodyCopy) Write(b []byte)(int,error){
	// 先记录
	bc.body.Write(b)

	// 再写在响应里
	return bc.ResponseWriter.Write(b)
}
// 全局中间件
// CopoyResp 复制请求的响应体
func CopoyResp(next http.HandlerFunc)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		// 初始化
		bc :=NewBodyCopy(w)

		next(bc,r)

		// 处理后
		fmt.Printf("")
	}
}