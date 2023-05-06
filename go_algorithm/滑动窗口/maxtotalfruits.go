package main

// 滑动窗口
func maxTotalFruits(fruits [][]int,startPos int,k int)int{
	left := 0 // 窗口左边界
	right := 0 // 窗口右边界
	n := len(fruits) // 存在水果位置的个数
	sum := 0 // 前缀和 
	ans := 0 // 答案
	step := func(left int,right int)int{
		if fruits[right][0] <= startPos{
			 // 只向左走的情况
			return startPos - fruits[left][0]
		}else if fruits[left][0]>=startPos{
			// 只向右走的情况
			return fruits[right][0] - startPos 
		}else{
			// 既向左又向右走的情况
			// l...pos...r
			// abs(pos-l) and abs(pos-r)
			// r - l 
			return min(abs(startPos -fruits[right][0]),abs(startPos-fruits[left][0]))+fruits[right][0] - fruits[left][0] 
		}

		// 每次固定窗口右边界
		for right <n {
			sum += fruits[right][1] // 说过个数前缀和
			// 移动左边界
			for left <= right && step(left,right) >k {
				sum  -= fruits[left][1] // 减去之前right在此位置加进去的值
				left++
			}
			ans = max(ans,sum)
			right++
		}
		return ans
	}

	func max(a,b int) int{
		if b>a{
			return b
		}
	}

	func min(a,b int)int{
		if a>b{
			return b
		}
	}


	func abs(x int)int{
		if x < 0 {
			return -x
		}
		return x
	}
	
}