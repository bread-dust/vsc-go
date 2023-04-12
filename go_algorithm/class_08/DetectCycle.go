package main

// 链表的入环节点
// https://leetcode-cn.com/problems/linked-list-cycle-ii/
func detectCycle(head *ListNode) *ListNode{
	if head == nil || head.Next == nil{
		return nil
	}

	slow := head.Next
	fast := head.Next.Next

	for slow !=fast {
		if fast.Next == nil|| fast.Next.Next == nil{
			return nil
		}	
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 有环
	fast = head
	for slow != fast{
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
