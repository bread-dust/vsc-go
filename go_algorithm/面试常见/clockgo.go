package main

import (
	"fmt"
	"time"
)

func g_ticker() {
	// Ticker 包含一个通道字段C，每隔时间段 d 就向该通道发送当时系统时间。
	// 它会调整时间间隔或者丢弃 tick 信息以适应反应慢的接收者。
	// 如果d <= 0会触发panic。关闭该 Ticker 可以释放相关资源。
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	ticker1 := time.NewTicker(5 *time.Second)
	defer ticker1.Stop()

	go func(t *time.Ticker){
		for {
			//每5s从chan.C中读取一次
			<-t.C
			fmt.Println("ticker:",time.Now().Format("2006-01-02 15:04:05"))
		}
	}(ticker1)
	time.Sleep(11*time.Second)
	fmt.Println("OK")
}

func g_timer(){
	fmt.Println("time:",time.Now().Format("2006-01-02 15:04:05"))	
	    // NewTimer 创建一个 Timer，它会在最少过去时间段 d 后到期，向其自身的 C 字段发送当时的时间
	timer1 := time.NewTimer(5 *time.Second)
	go func(t *time.Timer){
		times := 0
		for {
			<-t.C
			fmt.Println("timer:",time.Now().Format("2006-01-02 15:04:05"))
			 // 从t.C中获取数据，此时time.Timer定时器结束。如果想再次调用定时器，只能通过调用 Reset() 函数来执行
            // Reset 使 t 重新开始计时，（本方法返回后再）等待时间段 d 过去后到期。
            // 如果调用时 t 还在等待中会返回真；如果 t已经到期或者被停止了会返回假。
			times++
			// 调用reset 重发数据到chan C
			fmt.Println("调用reset重新设置定时器")
			t.Reset(2*time.Second)
			if times >3 {
				fmt.Println("停止调用")
				t.Stop()
			}
		}
	}(timer1)
	time.Sleep(11*time.Second*30)
	fmt.Println("结束时间",time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("ok")

}