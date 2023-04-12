package linkedlist

import "fmt"

// DoubleListNode is a node of double linked list
type DoubleListNode struct{
	Val  int
	Next *DoubleListNode
	Prev *DoubleListNode
}


func InitDoubleList() *DoubleListNode {
	return &DoubleListNode{
		Val:  0,
		Next: nil,
		Prev: nil,
	}
}

func (head *DoubleListNode) PrintDoubleLinkedList() {
	for head != nil {
		fmt.Printf("%d ",head.Val)
		head = head.Next
	}
	fmt.Println()
}

func (head *DoubleListNode) ReverseDoubleList() *DoubleListNode{
	var pre *DoubleListNode 
	var next *DoubleListNode
	for head != nil {
		next = head.Next
		head.Prev = next
		head.Next = pre
		pre = head
		head = next 
	}
	return pre
}

