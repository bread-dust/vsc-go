/*
@author:Deng.l.w
@version:1.20
@date:2023-02-18 12:25
@file:RandomArray.go
*/

package main

import (
	"math/rand"
	"time"
)


// 对数器
func generateRandomArray(n, v int) []int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	size := rand.Intn(n) + 1
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(v) + 1
	}
	return arr
}
