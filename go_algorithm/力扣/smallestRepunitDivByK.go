package main

func smallestRepunitDivByK(K int ) int{
	// 如果k是2或者5的倍数，那么一定不存在这样的数
	if K % 2 == 0 || K % 5 == 0{
		return -1
	}

	// cur 余数， ans 数字1的个数
	cur , ans := 0,1
	// 循环计算余数

	for {
		cur = (cur * 10 + 1) % K
		if cur == 0{
			return ans
		}
		ans++
	}

}