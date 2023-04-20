package main

import "fmt"

func main() {
	arr := []int{1, 3, 9, 5, 2, 5, 6, 2, 3}

	// fmt.Println(partition(arr, 2, len(arr)-1, 8))
	// fmt.Println(netherlands1(arr, 2, len(arr)-1, 8))
	fmt.Println(netherlands2(arr, 0, len(arr)-1))
	fmt.Println(quickSort2(arr, 0, len(arr)-1))
	fmt.Println(quickSort1(arr, 0, len(arr)-1))
	x := 3
	fmt.Println(kthMin(arr, 0, len(arr)-1, len(arr)-x))
}