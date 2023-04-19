package main

import (
	"fmt"
	"sync"

)

func GetRandom() {
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i:=0;i<5;i++{
			out<- i
		}
		close(out) //防止阻塞
	}()
	go func() {
		defer wg.Done()
		for o := range out{
			fmt.Println(o)
		}
	}()
	wg.Wait()
}

