package queue

type DoubleListNode struct {
	Val  int
	Next *DoubleListNode
	Prev *DoubleListNode
}

type QueueByDoubleLinkedList struct {
	head *DoubleListNode
	tail *DoubleListNode
	size int
}

// IsEmpty
func (qd *QueueByDoubleLinkedList) IsEmpty()bool{
	return qd.size==0
}

// Size
func (qd *QueueByDoubleLinkedList) Size()int{
	return qd.size
}

// EnqueueHead
func (qd *QueueByDoubleLinkedList) EnqueueHead(Val int){
	newNode := &DoubleListNode{
		Val:  Val,
		}

	if qd.head == nil{
		qd.head = newNode
		qd.tail = newNode
	}else{
		qd.head.Prev = newNode
		newNode.Next = qd.head
		qd.head = newNode
	}
	qd.size++
}

// EnqueueTail
func (qd *QueueByDoubleLinkedList) EnqueueTail(Val int){
	newNode := &DoubleListNode{
		Val:  Val,
		}

	if qd.head == nil{
		qd.head = newNode
		qd.tail = newNode
	}else{
		qd.tail.Next = newNode
		newNode.Prev = qd.tail
		qd.tail = newNode
	}
	qd.size++
}

// Dequeuehead
func (qd *QueueByDoubleLinkedList) DequeueHead()int{
	ans := qd.head.Val
	if qd.head == qd.tail{
		qd.head = nil
		qd.tail = nil
	}else{
		qd.head = qd.head.Next
		qd.head.Prev = nil
	}
	qd.size--
	return ans
}

// DequeueTail
func (qd *QueueByDoubleLinkedList) DequeueTail()int{
	ans := qd.tail.Val
	if qd.head == qd.tail{
		qd.head = nil
		qd.tail = nil
	}else{
		qd.tail = qd.tail.Prev
		qd.tail.Next = nil
	}
	qd.size--
	return ans
}
