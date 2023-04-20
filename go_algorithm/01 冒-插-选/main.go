package main

import "fmt"

func main() {
	testArr := []int{2, 1, 3, 6, 4, 6, 7, 3, 9}
	a := BubbleSort(testArr)
	b := insertSort(testArr)
	c := selectionSort(testArr)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}