/*
@author:Deng.l.w
@version:1.20
@date:2023-02-20 21:45
@file:KthLargest.go
*/

package main

import (
	"math/rand"
	"time"
)


func kthMin(arr []int, l, r, i int) int {
	if l == r { //l==r==i
		return arr[l]
	} else {
		rand.New(rand.NewSource(time.Now().UnixNano()))
		arr[l+rand.Intn(r-l+1)], arr[r] = arr[r], arr[l+rand.Intn(r-l+1)]
		x := partition3(arr, l, r)
		if x == i {
			return arr[x]
		} else if x > i {
			return kthMin(arr, l,x-1, i)
		} else {
			return kthMin(arr, x+1, r, i)
		}

	}
}
func partition3(arr []int, l, r int ) int{
	//小于等于区域右边界
	lessEqual := l - 1
	for i := l; i <= r; i++ {
		if arr[i] <= arr[r] {
			lessEqual++
			arr[lessEqual], arr[i] = arr[i], arr[lessEqual]
		}
	}
	return lessEqual
}
