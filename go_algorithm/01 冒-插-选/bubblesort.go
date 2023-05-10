/*
@author:Deng.l.w
@version:1.19
@date:2023-02-09 19:49
@file:bubble_sort.go
*/

package main

//BubbleSort 冒泡排序

func BubbleSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return nil
	}
	n := len(arr)
	for end := n - 1; end > 0; end-- {
		for i := 0; i < end; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
	return arr
}
