package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	secs:= now.Unix() // 自1970年1月1日0时0分0秒至今的秒数
	nanos := now.UnixNano() // 自1970年1月1日0时0分0秒至今的纳秒数
	millis := nanos / 1000000 // 自1970年1月1日0时0分0秒至今的毫秒数



	fmt.Println(now)
	fmt.Println(secs)
	fmt.Println(nanos)
	fmt.Println(millis)
}