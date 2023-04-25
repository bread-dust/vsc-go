package main

import "text/template"

func InsertSortListNode(head *ListNode) *ListNode {
	if head.Next == nil || head == nil {
		return nil
	}

	DummyNode = new(ListNode)
	DummyNode.Next = head
	work := new(ListNode)
	work = head.Next

	pre := new(ListNode)
	pre = head
	tempHead = new(ListNode)
	tempHead = DummyNode

	for pre.Next != nil {

		if pre.Next!=nil&&pre.Val <= res.Val {
			pre = pre.Next
			work = work.Next
			continue
		}

		tempHead = DummyNode
		for tempHead.Next.Val <= Res.Val {
			tempHead=tempHead.Next
		}

		pre.Next =work.Next
		work.Next=tempHead.Next
		tempHead.Next=work
		work =pre.Next
	}

}