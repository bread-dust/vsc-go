package middleware

import (
	"core/helper"
	"net/http"
	"strconv"
)

type AuthMiddleware struct {
}

// NewAuthMiddleware create a new auth middleware
func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

// Handle handle the middleware
func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		// 从header中获取Authorization
		auth := r.Header.Get("Authorization")
		if auth == ""{ // 未授权
			w.WriteHeader(http.StatusUnauthorized) // 401
			w.Write([]byte("Unauthorized")) // 未授权
			return
		}

		// 解析token
		uc,err:=helper.AnalyzeToken(auth)
		if err!=nil{ //无法正确解析
			w.WriteHeader(http.StatusUnauthorized) // 401
			w.Write([]byte(err.Error())) // 未授权
			return 
		}

		// 将用户信息写入header
		r.Header.Set("userId",strconv.Itoa(int(uc.Id)))
		r.Header.Set("UserIdentity",uc.Identity)
		r.Header.Set("UserName",uc.Name)
		
		// Passthrough to next handler if need
		next(w, r)
	}
}
