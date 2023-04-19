package main

import (
	"fmt"
	"sync"
)

func Goroutine_alternately() {
	letter, number := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{} // 等待组
	go func() {
		i := 1
		for {
			select {
			case <-number:
				print(i)
				i++
				print(i)
				i++
				letter <- true
			}
		}
	}()
	wait.Add(1) 

	go func(wait *sync.WaitGroup) {
		j := 'A'
		for{
			select{
			case <-letter:
				if j >= 'Z'{
					fmt.Println("")
					wait.Done()
					return 
				}
				print(string(j))
				j++
				print(string(j))
				j++
				number <- true
			}
		}
	}(&wait)

	number <- true // 程序开始
	wait.Wait() // letter阻塞在此，等待结束
}