package main

import "fmt"

func atn() {
	// 113 , 23  113 + 23 = 136
	// 342,465 
	//
	listNode1 := &ListNode{Val: 2}
	listNode2 := &ListNode{Val: 4}
	listNode3 := &ListNode{Val: 3}

	listNode4 := &ListNode{Val: 5}
	listNode5 := &ListNode{Val: 6}
	listNode6 := &ListNode{Val: 4}
	// listNode7 := &ListNode{Val: 8}

	listNode1.Next = listNode2
	listNode2.Next = listNode3

	listNode4.Next = listNode5
	listNode5.Next = listNode6
	// listNode6.Next = listNode7


	res := addTwoNumbers(listNode1, listNode4)
	for res != nil {
		fmt.Printf("%d ", res.Val)
		res = res.Next
	}
}
func main() {
	atn()
}

描述：
给你一个链表的头节点head,每k个节点一组进行翻转，请你返回修改后的链表。
k是一个正整数，它的值小于或等于链表的长度。如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。
你不能更改节点中的值，只能更改节点本身。

思路：
思路1： 容器：数组里交换
思路2：