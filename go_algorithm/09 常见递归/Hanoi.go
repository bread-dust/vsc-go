package main

import "fmt"

// 汉诺塔的移动问题

func Hanoi(n int){
	if n>0{
		process3(n,"left","right","mid")
	}
}

func process3(n int,from,to,other string){
	if n==1{
		fmt.Println("move",1,"from",from,"to",to)
	}else{
		// n - 1 移到 中间
		process3(n-1,from,other,to)
		// n 移到右边
		fmt.Println("move",n,"from",from,"to",to)
		// n-1 移到右边
		process3(n-1,other,to,from)
	}
}