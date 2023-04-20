package heapsort

import "heap"


// 堆排序
func SortArray(arr []int)[]int{
	if arr==nil||len(arr)<2{
		return arr
	}
	n:=len(arr)

	maxHeap := heap.NewHeap()
	// 建立大根堆
	for i:=0;i<n;i++{
		maxHeap.Push(arr[i]) // 堆插入
	}

	maxHeap.Swap(0,n-1) // 交换堆顶和最后一个元素

	for heapSize:=n-1;heapSize>0;heapSize--{
		maxHeap.Heapify(0,heapSize)
		maxHeap.Swap(0,heapSize-1)
		
	}
	return arr
}



