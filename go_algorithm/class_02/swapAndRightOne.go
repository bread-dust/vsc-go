/*
@author:Deng.l.w
@version:1.20
@date:2023-02-17 23:14
@file:swapAndRightOne.go
*/

package main
// swapAndRightOne 提取最右侧的1
func swapAndRightOne(a int) int{
	/*
		1. 最右侧的1的右侧一定为0或已是末尾。
		2. 1在末尾，1取反为0，+1为1，左边不受影响，和原值取合集
		2. 末尾0 经取反为1，+1变为0，向高位进1
		3. 最右侧的1取反变为0，收到低位0的进1，变为1
		4. 最右侧的1的左侧均不受影响，仅取反
		5. 结果和原值取合集
	*/
	for a != 0 {
		/*
			a:			0010 0101 00
			rightOne:	0000 0001 00 -- 4
			a^rightOne: 0010 0100 00

			a:			0010 0100 00
			rightOne: 	0000 0100 00 -- 16
			a^rightone: 0010 000 00
		*/
		rightOne := a & (-a + 1)
		a ^= rightOne
	}
	return a
}
