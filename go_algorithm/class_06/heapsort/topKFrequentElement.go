package heapsort

import (
	goheap "container/heap"
)

type Record struct{
	number int 
	times int 
}

type RecordTimesMinHeap []Record

// 实现heap.Interface接口
// Len()实现了求长
func (h RecordTimesMinHeap) Len() int { return len(h) } 
// Less()实现了比较,谁小谁放在顶部，小根堆
func (h RecordTimesMinHeap) Less(i, j int) bool { return h[i].times < h[j].times }
// Swap()实现了交换
func (h RecordTimesMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
// Push()实现了加入
func (h *RecordTimesMinHeap) Push(x interface{}) {
	*h = append(*h, x.(Record))
}
// Pop()实现了弹出
func (h *RecordTimesMinHeap) Pop() interface{} {
	old := *h // 拷贝一份
	n := len(old) // 拷贝的长度
	x := old[n-1] // 拷贝的最后一个元素
	*h = old[0 : n-1] // 原来的切片去掉最后一个元素
	return x
}


// 求出现次数前k高的元素

func TopKFrequent(nums []int,k int) [] int{
	numberTimes:= make(map[int]int)  // 次数哈希表,key是数字，value是出现次数
	n := len(nums) // 数组长度
	// 给每一种数建立词频记录
	// O(n)
	for i:=0;i<n;i++{
		if _,ok :=numberTimes[nums[i]];!ok{
			numberTimes[nums[i]] =0 // 如果没有记录
		}
		numberTimes[nums[i]]++ // 词频加1
	}

	// 建立词频小根堆
	h:=&RecordTimesMinHeap{}
	goheap.Init(h)

	for number,times:=range numberTimes{
		if h.Len()<k{ // 堆没满
			goheap.Push(h,Record{number,times})
		}else{ 
			//堆满了
			// 看当前的数和对应的词频能不能秒掉 门槛
			top := goheap.Pop(h)

			if top.(Record).times>=times {
				// 不加入当前数和其词频
				// 把门槛加回去
				goheap.Push(h,top)
			}else{ //当前数的词频> 门槛
				//加入当前数和词频
				//不加门槛回去
				goheap.Push(h,Record{number,times})
				
			}
		}
	}

	ans:=make([]int,k) // 结果数组
	for i:=0;h.Len()>0;i++{
		ans[i]=goheap.Pop(h).(Record).number
	}
	return ans
}

