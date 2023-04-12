package stack

type ListNode struct {
	Val  int
	Next *ListNode
}


type LinkedListStack struct {
	head *ListNode
	size int
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		head: nil,
		size: 0,
	}
}

func (lsk *LinkedListStack) IsEmpty()bool{
	return lsk.size == 0
}

func (lsk *LinkedListStack) Size()int{
	return lsk.size
}

func (lsk *LinkedListStack) Push(Val int){
	oldHead := lsk.head
	lsk.head = &ListNode{
		Val:  Val,
	}
	lsk.head.Next = oldHead
	lsk.size++
}

func (lsk *LinkedListStack) Pop()int{
	ans := lsk.head.Val
	lsk.head = lsk.head.Next
	lsk.size--
	return ans
}

func (lsk *LinkedListStack) Peek()int{
	return lsk.head.Val
}


