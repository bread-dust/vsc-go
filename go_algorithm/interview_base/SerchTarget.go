package main

import (
	"context"
	"fmt"
	"time"
)

func SearchTarget() {
	timer := time.NewTicker(time.Second*5) //定时器
	ctx,cancel:=context.WithCancel(context.Background()) // 上下文对象，控制gorountine工作
	resultChan:=make(chan bool) // 结果通道

	data := []int{1,2,3,4,5,6,7,8,9}
	datalen:= len(data)
	size := 2
	tartget := 4
	for i:=0;i<datalen;i+=size{ // 分片大小为size
		end := i +size
		if end >= datalen -1{
			end = datalen-1
		}
	go Work(ctx,data[i:end],tartget,resultChan)
	}

	select {
	case <-timer.C: // 超时
		fmt.Println("tiemout,notfound")
		cancel() // 给goroutine发送结束信号
	case <-resultChan: // 结果通道有值
		fmt.Println("found") // 找到
		cancel() // 给goroutine发送结束信号
	}
	time.Sleep(time.Second*2)
}

func Work(ctx context.Context,data []int,target int,resultChan chan bool) {
	for _,v:=range data{
		select {
			case <-ctx.Done(): // ctx停止goroutnie工作
				fmt.Println("task cancel")
				return
			default:
		}

		fmt.Printf("v:%d",v) // 打印遍历值
		time.Sleep(time.Second*1) // 等待1s
		if target == v{
			resultChan <- true // 结果通道存放true
			return
		}
	}
}