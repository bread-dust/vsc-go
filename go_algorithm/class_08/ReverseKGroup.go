package main

// 每组K个节点之间逆序
// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/

func reverseKGroup(head *ListNode, k int ) *ListNode{
	start := head // start->a
	end := getKGroupEnd(start,k)
	//abc def g
	if end == nil{ // 链表不够k个
		return head 
	}

	head = end // 新的头部 head -> c
	reverse(start,end) // cba->def->g
	lastEnd := start  // lastEnd -> a
	for lastEnd.Next!=nil{ // a.Next !=null
		start = lastEnd.Next // start -> d
		end = getKGroupEnd(start,k) //end -> f
		if end!=nil{
			return head  
		}
		reverse(start,end) //cba->fed->g
		lastEnd.Next=end // a.Next=f
		lastEnd = start // lastEnd -> start -> d
	}

	return head
} 

// getKGroupEnd 获得第k组的结尾
// 如果对当前组不够，返回nil
func getKGroupEnd(start *ListNode,k int) *ListNode{
	for i := k-1;i!=0&&start != nil;i--{
		start= start.Next
	}
	return start
}

func reverse(start *ListNode,end *ListNode) {
	newstart := end.Next  // end->1,下一段的起始节点
	var pre *ListNode
	var next *ListNode
	cur := start // 当前操作的节点
	for cur!=end {
		next = cur.Next // 先存下一步要操作节点的下标
		cur.Next = pre // 改变方向
		pre = cur // 当前节点的前一节点
		cur = next // 当前节点向后移动
	}
	start.Next = newstart //当前段的尾结点，连接到下一段的起始节点
	 
}