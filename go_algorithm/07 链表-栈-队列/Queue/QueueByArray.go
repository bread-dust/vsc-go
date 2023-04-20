package queue

type ArrayQueue struct {
	data []int
}

func NewArrayQueue() *ArrayQueue{
	return &ArrayQueue{
		data: make([]int, 0),
	}
}


func (aq *ArrayQueue) IsEmpty()bool{
	return len(aq.data) == 0
}

func (aq *ArrayQueue) Size() int{
	return len(aq.data)
}

func (aq *ArrayQueue) Enqueue(Val int){
	aq.data = append(aq.data, Val)
}

func (aq *ArrayQueue) Dequeue()int{
	ans := aq.data[0]
	aq.data = aq.data[1:]
	return ans
}

func (aq *ArrayQueue) Peek()int{
	return aq.data[0]
}
