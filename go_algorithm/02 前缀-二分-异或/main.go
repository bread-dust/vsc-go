package main

import "fmt"

func main() {
	arr := []int{10, 3, 0, -9, 2, -22, 54, 23, 7}
	fmt.Println(moreEqualMostRight(arr, 3)) 
	fmt.Println(moreEqualMostLeft(arr,4))
	fmt.Println(moreEqualMostRight(arr, 4))
	fmt.Println(exist(arr, 3))
	fmt.Println(oddTimesOneKind(arr))
	fmt.Println(oddTimestwoKind(arr))
	fmt.Println(getSum(preSumArray(arr), 2, 3))
	fmt.Println(swapAndRightOne(3))
}