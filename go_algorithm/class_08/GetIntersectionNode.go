package main
// 一个链表的首次相交节点
// https://leetcode-cn.com/problems/intersection-of-two-linked-lists/

func getIntersectionNode(head1,head2 *ListNode)*ListNode{
	if head1 == nil|| head2 == nil{
		return nil
	}

	cur1 := head1
	cur2 := head2 
	n := 0
	for cur1.Next !=nil{
		n++
		cur1 = cur1.Next
	}

	for cur2.Next != nil{
		n--
		cur2 = cur2.Next
	}

	if cur1 != cur2{  // 如果两个链表没有交点，那么最后一个节点一定不相同
		return nil
	}

	cur1 = ternary(n>0,head1,head2).(*ListNode) // 使得cur1指向较长的链表
	cur2 = ternary(cur1==head1,head2,head1).(*ListNode) // 使得cur2指向较短的链表

	if n<0{n= -n}
	for n!=0{ // 长链表走差值步
		n--
		cur1=cur1.Next
	}
	for cur1!=cur2{
		cur1=cur1.Next
		cur2=cur2.Next
	}	
	return cur1
}