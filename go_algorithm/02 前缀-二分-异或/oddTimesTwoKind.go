package main


func oddTimestwoKind(arr []int) (int, int) {
	eor := 0
	for i := 0; i < len(arr); i++ {
		eor ^= arr[i]
	}
	/*
	eor = a^b
	 */
	rightOne := eor & (-eor + 1)
	// rightOne 是 a^b 最右侧 1的位
	eorr := 0
	for j := 0; j < len(arr); j++ {
		if (rightOne&arr[j]) != 0 {
			// 找出rightOne位 为 1  的数
			eorr ^= arr[j]
			// a ^ b ^ b
		}
	}
	 a := eorr
	 b := eor ^ eorr
	 return a,b
}