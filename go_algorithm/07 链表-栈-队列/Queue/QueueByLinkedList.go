package queue

type ListNode struct{
	Val int
	Next *ListNode
}

type LinkedListQueue struct{
	head *ListNode
	tail *ListNode
	size int
}

func (lq *LinkedListQueue) Size()int{
	return lq.size
}

func (lq *LinkedListQueue) IsEmpty()bool{
	return lq.size == 0
}

func (lq *LinkedListQueue) Peek()int{
	return lq.head.Val
}

func (lq *LinkedListQueue) Enqueue(Val int){
	newNode := &ListNode{
		Val:  Val,
	}
	if lq.head == nil{
		lq.head = newNode
		lq.tail = newNode
	}else {
		lq.tail.Next = newNode
		lq.tail = newNode
	}
	lq.size++
}

func (lq *LinkedListQueue) Dequeue()int{
	ans := lq.head.Val
	if lq.head == lq.tail{
		lq.head = nil
		lq.tail = nil
	}else {
		lq.head = lq.head.Next
	}
	lq.size--
	return ans
}
