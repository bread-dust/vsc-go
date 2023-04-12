package linkedlist

import "fmt"

// 单链表节点
type SingledListNode struct {
	Val  int       // value of node
	Next *SingledListNode // pointer to next node
}

func InitList() *SingledListNode {
	return &SingledListNode{
		Val:  0,
		Next: nil,
	}
}

func (head *SingledListNode) PrintSingledLinkedList() {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

// ReverseList 单链表翻转
func (head *SingledListNode) ReverseList() *SingledListNode {
	var pre *SingledListNode // previous node
	var next *SingledListNode // next node
	for head != nil { // loop until h is nil
		next = head.Next // next游标指针暂存下一跳节点
		head.Next = pre // 当前节点指向pre游标指针
		pre = head	// pre游标指针指向当前节点
		head = next // head游标指针指向下一跳节点
}
	return pre
}
