/*
@author:Deng.l.w
@version:1.20
@date:2023-02-10 10:19
@file:Presum.go
*/

package main


func preSumArray(arr []int) []int {
	n := len(arr)
	sum := make([]int, n)
	sum[0] = arr[0]
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + arr[i]
	}
	return sum
}

func getSum(arr []int, i int, j int) int {
	if i == 0 {
		return arr[j]
	} else {
		return arr[j] - arr[i-1]
	}

}
