/*
@author:Deng.l.w
@version:1.19
@date:2023-02-09 23:07
@file:insert_sort.go
*/

package main

// insertSort 插入排序
// 小范围先排好，在逐渐扩大范围
func insertSort(arr []int) []int{
	if arr == nil || len(arr) < 2 {
		return nil
	}
	n := len(arr)

	for i := 1; i < n-1; i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
	return arr
}
