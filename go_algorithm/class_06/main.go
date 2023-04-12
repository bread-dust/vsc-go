package main

import (
	"fmt"
	"heapsort"
)
func main() {
	
	nums := []int{5,2,3,1}
	arrs := []int{0,0,1,1,2,5}

	// 1,2,3,5
	// 001125
	fmt.Println(heapsort.SortArray(nums))
	fmt.Println(heapsort.SortArray(arrs))

	fmt.Println(heapsort.TopKFrequent(arrs,2))

}