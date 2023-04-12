/*
@author:Deng.l.w
@version:1.20
@date:2023-02-19 16:54
@file:Netherlands.go
*/

package main

func netherlands1(arr []int, l, r, pivot int) {
	less := l - 1
	more := r + 1
	for i := l; i <= r; {
		if arr[i] < arr[pivot] {
			less++
			arr[less+1], arr[i] = arr[i], arr[less+1]
			i++
		} else if arr[i] > arr[pivot] {
			arr[more-1], arr[i] = arr[i], arr[more-1]
			more--
		} else {
			i++
		}
	}
}

// arr[r] 作为划分值
// 返回值是数组，左右边界的下标
func netherlands2(arr []int, l, r int) []int {
	less := l - 1
	more := r
	for i := l; i <= more; {
		if arr[i] < arr[more] {
			less++
			arr[less], arr[i] = arr[i], arr[less]
			i++
		} else if arr[i] > arr[r] {
			arr[more-1], arr[i] = arr[i], arr[more-1]
			more--
		} else {
			i++
		}
	}
	arr[more], arr[r] = arr[r], arr[more] // 将划分值放到中间
	return []int{less + 1, more}
}
