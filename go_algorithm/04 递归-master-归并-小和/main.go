package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 6, 3, 2, 8, 2, 6,10,4}
	fmt.Println(getMax(arr, 2, 6))
	fmt.Println(sortArray(arr))
	fmt.Println(sortArrayUnrecursive(arr))
	fmt.Println(smallSum(arr))
}