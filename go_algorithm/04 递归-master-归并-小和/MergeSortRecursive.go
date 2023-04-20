/*
@author:Deng.l.w
@version:1.20
@date:2023-02-19 9:58
@file:MergeSortRecursive.go
*/

package main


func sortArray(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	n := len(arr)
	help := make([]int, n)
	mergeSort(arr, help, 0, n-1)
	return arr 
}

func mergeSort(arr, help []int, l, r int) {
	// base
	if l >= r {
		return
	}
	// 中点
	m := l + (r-l)/2
	// 左部分有序
	mergeSort(arr, help, l, m)
	// 右部分有序
	mergeSort(arr, help, m+1, r)
	//整合
	merge(arr, help, l, m, r)
}

func merge(arr, help []int, l, m, r int) {
	p1 := l
	p2 := m + 1
	i := l
	for ; p1 <= m && p2 <= r; i++ { //p1,p2不越界
		if arr[p1] > arr[p2] { //p1>p2
			help[i] = arr[p2]
			p2++ //p2移到下一个
		} else {
			help[i] = arr[p1]
			p1++ //p1移到下一个
		}
	}
	for ; p1 <= m; i++ { //如果p2先越界，p1不越界，把p1剩下的都放到help里
		help[i] = arr[p1]
		p1++
	}
	for ; p2 <= r; i++ { // 如果p1先越界，p2不越界，把p2剩下的都放到help里
		help[i] = arr[p2]
		p2++
	}
	for i := l; i <= r; i++ { //把help里的数放回arr
		arr[i] = help[i]
	}
}
