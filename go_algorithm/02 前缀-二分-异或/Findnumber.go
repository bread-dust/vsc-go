/*
@author:Deng.l.w
@version:1.20
@date:2023-02-10 10:46
@file:Findnumber.go
*/

package main


// arr 是有序数组，使用二分法
func exist(arr []int, num int) bool {
	if arr == nil || len(arr) == 0 {
		return false
	}
	l := 0
	r := len(arr) - 1
	m := 0
	for l <= r {
		m = l + (r-l)/2
		if arr[m] == num {
			return true
		} else if arr[m] > num {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return false
}
