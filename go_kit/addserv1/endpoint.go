package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"

)

// endpont

// 一个 endpoint 表示对外提供的一个方法
// 1.3. 请求和响应
type SumRequest struct{
	A int `json:"a"`
	B int `json:"b"`
}

type SumReponse struct{
	V int `josn:"v"`
	Err string `json:"err,omitempty"`
}

type ConcatRequest struct{
	A string `json:"a"`
	B string `json:"b"`
}

type ConcatReponse struct{
	V string `josn:"v"`
	Err string `json:"err,omitempty"`
}

// 2. EndPoint 借助适配器 将方法->endponit 
func makeSumEndponit(srv Addservice) endpoint.Endpoint{
	return func(ctx context.Context,request interface{})(interface{},error){
		req := request.(SumRequest) // 断言请求类型， endpoint 只接受 SumRequest请求
		v,err := srv.Sum(ctx,req.A,req.B) // 进行方法调用，取得结果
		if err!=nil{
			return SumReponse{V:v,Err: err.Error()},err 
		}
		return SumReponse{V:v},nil // 将结果封装在 SumResponse 返回暴露给上层
	}
}

func makeConcatEndponit(srv Addservice) endpoint.Endpoint{
	return func(ctx context.Context,request interface{})(interface{},error){
		req := request.(ConcatRequest)
		v,err := srv.Concat(ctx,req.A,req.B) // 方法调用
		if err!=nil{
			return ConcatReponse{V:v,Err: err.Error()},err
		}
		return ConcatReponse{V:v},nil

	}
}

