/*
@author:Deng.l.w
@version:1.20
@date:2023-02-17 22:40
@file:bsNearLeft.go
*/

package main
/*
	有序数组不大于val
*/
func moreEqualMostRight(arr []int, val int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	l := 0
	r := len(arr) - 1
	ans := -1
	for l <= r {
		m := l + (r-l)>>1
		if arr[m] <= val {
			ans = m
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return ans
}
