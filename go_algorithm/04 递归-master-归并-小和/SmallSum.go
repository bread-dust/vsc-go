/*
@author:Deng.l.w
@version:1.20
@date:2023-02-19 15:52
@file:SmallSum.go
*/

package main

// 暴力解法
// func right(arr []int) int {
// 	ans := 0
// 	length := len(arr)
// 	for i := 1; i < length; i++ {
// 		for j := 0; j < i; j++ {
// 			if arr[j] < arr[i] {
// 				ans += arr[j]
// 			}
// 		}
// 	}
// 	return ans
// }

func smallSum(arr []int) int {
	if arr == nil && len(arr) < 2 {
		return 0
	}
	n := len(arr)
	help := make([]int, n)
	return process(arr, help, 0, n-1)
}

func process(arr, help []int, l, r int) int {
	if l >= r {
		return 0
	}
	m := l + (r-l)/2
	left := process(arr, help, l, m)
	right := process(arr, help, m+1, r)
	merge := merge1(arr, help, l, m, r)
	return left + right + merge
}

func merge1(arr, help []int, l, m, r int) int {
	ans := 0
	p1 := l
	p2 := m + 1
	i := l
	for ; p1 <= m && p2 <= r; i++ {
		if arr[p1] < arr[p2] {
			ans += (r - p2 + 1) * arr[p1]
			help[i] = arr[p1]
			p1++
		} else {
			help[i] = arr[p2]
			p2++
		}
	}
	for ; p1 <= m; i++ {
		help[i] = arr[p1]
		p1++
	}
	for ; p2 <= r; i++ {
		help[i] = arr[p2]
		p2++
	}
	for i = l; i <= r; i++ {
		arr[i] = help[i]
	}
	return ans
}
