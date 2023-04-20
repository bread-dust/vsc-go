package queue

import "fmt"

type QueueByRingArray struct {
	data     []int
	capacity int // 容量
	size     int // 已经存放的元素个数
	start    int // 管理pop 的下标
	end      int // 管理push的下标
}

func NewQueueByRingArray(capacity int) *QueueByRingArray {
	return &QueueByRingArray{
		data:     make([]int, capacity),
		capacity: capacity,
		size:     0,
		start:    0,
		end:      0,
	}
}
func (q *QueueByRingArray) IsEmpty() bool {
	return q.size == 0
}

func (q *QueueByRingArray) IsFull() bool {
	return q.size == q.capacity
}

func (q *QueueByRingArray) Size() int {
	return q.size
}

func (q *QueueByRingArray) Enqueue(Val int) {
	if q.IsFull() {
		fmt.Println("已满")
	}else{
		q.data[q.end] = Val
		q.end = (q.end +1) %q.capacity
		}	
	q.size++
}

func (q *QueueByRingArray) Dequeue()int{
	if q.IsEmpty(){
		fmt.Println("已空")
		return -1
	}else{
		ans := q.data[q.start]
		q.start = (q.start +1) %q.capacity
		q.size--
		return ans
	}
}