package main

import "fmt"

func main() {
	a := []int{1, 2,3}
	result := subsets(a)
	for _, v := range result {
		fmt.Println(v)
	}

	result2 :=permute(a)
	fmt.Println(result2)
	
	n:=3 // 三层
	Hanoi(n)
}