package main

import (
	"context"
	"errors"
)

// service 层
// 存放所有业务逻辑
// 1.1 业务逻辑抽象为接口

type Addservice interface{  //服务集合的抽象模块接口
	Sum(ctx context.Context,a,b int)(int,error)
	Concat(ctx context.Context,a,b string)(string,error)
}

// 1.2 Addservice接口模块的具体实现
type addService struct{
	// 可存放数据库
	// 可扩展各种字段

}
var (
	ErrEmptyString = errors.New("两个字符串都是空字符串")
)

// NewService ,Addservice 的构造函数,生成一个服务集合模块
func NewService()Addservice{
	return addService{}
}

// Sum 返回和，定义服务
func (addService)Sum(ctx context.Context,a,b int)(int,error){
	return a+b,nil
}

// Concat 拼接字符串，定义服务
func (addService)Concat(ctx context.Context,a,b string)(string,error){
	if a == "" && b == ""{
		return "", ErrEmptyString
	}
	return a+b,nil
	
}

