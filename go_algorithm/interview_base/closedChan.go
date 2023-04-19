package main

import "fmt"

func WriteClosedChan() {
	c := make(chan int, 3)
	close(c)
	c <- 1
}

func ReadClosedChan() {
	fmt.Println("以下是数值的chan")
	ci := make(chan int,2)
	ci <-1
	close(ci)
	num,ok := <- ci
    fmt.Printf("读chan的协程结束，num=%v， ok=%v\n",num,ok)
    num1,ok1 := <-ci
    fmt.Printf("再读chan的协程结束，num=%v， ok=%v\n",num1,ok1)

	fmt.Println("以下是字符串chan")
    cs := make(chan string,2)
    cs <- "aaa"
    close(cs)
    str,ok := <- cs
    fmt.Printf("读chan的协程结束，str=%v， ok=%v\n",str,ok)
    str1,ok1 := <-cs
    fmt.Printf("再读chan的协程结束，str=%v， ok=%v\n",str1,ok1)
    
	fmt.Println("以下是结构体chan")
    type MyStruct struct{
        Name string
    }
    cstruct := make(chan MyStruct,2)
    cstruct <- MyStruct{Name: "haha"}
    close(cstruct)
    stru,ok := <- cstruct
    fmt.Printf("读chan的协程结束，stru=%v， ok=%v\n",stru,ok)
    stru1,ok1 := <-cs
    fmt.Printf("再读chan的协程结束，stru=%v， ok=%v\n",stru1,ok1)
}