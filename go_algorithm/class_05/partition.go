/*
@author:Deng.l.w
@version:1.20
@date:2023-02-19 16:39
@file:partition.go
*/

package main


// arr[l...r],给定pivot 为划分值
func partition(arr []int, l, r, pivot int) {
	//小于等于区域右边界
	lessEqual := l - 1
	for i := l; i <= r; i++ {
		if arr[i] <= pivot {
			lessEqual++
			arr[lessEqual], arr[i] = arr[i], arr[lessEqual]
		}
	}
}
