/*
@author:Deng.l.w
@version:1.20
@date:2023-02-17 23:36
@file:oddTimesOneKind.go
*/

package main



func oddTimesOneKind(arr []int) int{

	eor := 0
	for i := 0; i < len(arr); i++ {
		eor = eor ^ arr[i]
	}
	return eor
}
