/*
@author:Deng.l.w
@version:1.19
@date:2023-02-09 18:35
@file:select_sort.go
*/

package main

// selectionSort 选择排序
// 先假定选择好最小值下标，再交换
func selectionSort(arr []int) []int{
	if arr == nil || len(arr) < 2 {
		return nil
	}
	// 未满足条件 不排序
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i // 假定最小值下标
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}
