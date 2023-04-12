/*
@author:Deng.l.w
@version:1.19
@date:2023-02-09 19:50
@file:pracetice.go
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	m := 2 ^ 3
	fmt.Printf("%d", m)
	arr := []int{1, 2, 3}
	/*
		出现偶次数的数互相异或为0，只剩奇数次的数变为0^自身
	*/
	sort.SliceIsSorted(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}
