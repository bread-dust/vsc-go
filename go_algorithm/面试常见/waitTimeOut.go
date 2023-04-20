package main

import (
    "fmt"
    "sync"
    "time"
)

func WaitTime() {
    wg := sync.WaitGroup{}
    c := make(chan struct{})
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(num int, close <-chan struct{}) {
            defer wg.Done() // 最后结束
            <-close
            fmt.Println(num)
        }(i, c)
    }

    if WaitTimeout(&wg, time.Second*5) {
		// 已经超时
        close(c)
        fmt.Println("timeout exit")
    }
	// 没超时，等待10s
    time.Sleep(time.Second * 10)
}

func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
    // 要求手写代码
    // 要求sync.WaitGroup支持timeout功能
    // 如果timeout到了超时时间返回true
    // 如果WaitGroup自然结束返回false
	ch := make(chan bool,1) // 保存是否超时的信号

	// 5s 到了超时时间后将通道放进true，表示已超时
	go time.AfterFunc(timeout, func() {
		ch <- true
	})

	// 正常wait,将通道放进false，表示没超时
	go func(){
		wg.Wait()
		ch<-false
	}()

	return <-ch
}