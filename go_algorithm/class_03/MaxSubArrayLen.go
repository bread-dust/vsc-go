/*
@author:Deng.l.w
@version:1.20
@date:2023-02-18 11:22
@file:MaxSubArrayLen.go
*/

package main


func maxSubArrayLen(nums []int, k int) int {
	//ans 最大长度
	//sum 前缀和
	ans := 0
	sum := 0
	// key 前缀和，value 前缀和最早出现在0-value上
	sumMap := map[int]int{
		0: -1, // 一个数也没有，就存在前缀和0
	}
	for i, num := range nums {
		sum += num
		pre, ok := sumMap[sum-k] // sum-k 最早出现位置 pre
		// (0...pre)..i
		if ok && ans < i-pre {
			ans = i - pre
		}
		// 前缀和：下标 加入sumMap
		if _, ok := sumMap[sum]; !ok {
			sumMap[sum] = i
		}
	}
	return ans
}
