/*
@author:Deng.l.w
@version:1.20
@date:2023-02-20 20:38
@file:quickSort.go
*/

package main

import (
	"math/rand"
	"time"
)


func quickSort1(arr []int, l, r int)[]int {
	if l > r {
		rand.New(rand.NewSource(time.Now().UnixNano()))
		arr[l+rand.Intn(r-l+1)], arr[r] = arr[r], arr[l+rand.Intn(r-l+1)]
		mid := partition2(arr, l, r)
		quickSort1(arr, l, mid-1)
		quickSort1(arr, mid+1, r)
	}
	return arr

}

// 搞定一批相同的数
func quickSort2(arr []int, l, r int) []int{
	if l < r {
		x := netherlands2(arr, l, r)
		quickSort2(arr, l, x[0]-1)
		quickSort2(arr, x[1]+1, r)
	}
	return arr
}

// 默认以arr[r]为基准
func partition2(arr []int, l, r int) int{
	//小于等于区域右边界
	lessEqual := l - 1
	for i := l; i <= r; i++ {
		if arr[i] <= arr[r] {
			lessEqual++
			arr[lessEqual], arr[i] = arr[i], arr[lessEqual]
		}
		// 循环结束a[lessEqual] == arr[r]
	}
	return lessEqual
}