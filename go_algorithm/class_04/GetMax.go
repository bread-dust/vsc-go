/*
@author:Deng.l.w
@version:1.20
@date:2023-02-18 16:25
@file:GetMax.go
*/

package main


// 递归实现
func getMax(arr []int, l, r int) int {
	// basecase
	if l == r {
		return arr[l]
	}
	// l<r ,m是中点
	// l...m......r
	//l...m
	//      m+1...r
	m := (l + r) / 2 // l +(r-l)/2 q
	leftMax := getMax(arr, l, m)
	rightMax := getMax(arr, m+1, r)

	if leftMax >= rightMax {
		return leftMax
	} else {
		return rightMax
	}
}

/*
master公式qui递归
T(n)= a * T(N/b) + O(N ^d )
log(b,a)<d, O(N ^ d)
log(a,b)>d,O(N ^ log(b,a))
log(b,a) == d, O( ^ d * logN)
*/
