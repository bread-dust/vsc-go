/*
@author:Deng.l.w
@version:1.20
@date:2023-02-17 22:39
@file:bsNearRight.go
*/

package main


func moreEqualMostLeft(arr []int, val int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	l := 0
	r := len(arr) - 1
	ans := -1
	for l <= r {
		m := l + (r-1)>>1
		if arr[m] >= val {
			ans = m // m的极限  -> val
			r = m - 1 // r的极限 -> val-1
		} else {
			l = m + 1 // l 的极限 -> val+1
		}
	}
	return ans
}


