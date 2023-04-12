package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 6, 3, 2, 6, 2, 6, 10}
	hushMap()
	fmt.Println(maxSubArrayLen(arr, 10))
}