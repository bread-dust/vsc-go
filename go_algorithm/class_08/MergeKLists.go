package main

// 合并所有的有序链表
// https://leetcode-cn.com/problems/merge-k-sorted-lists/
func mergeKLists(lists []*ListNode) *ListNode{

	if len(lists) == 0{
		return nil
	}

	n := len(lists)
	h := &NodeHeap{}
	heap.Init(h)

	for i:=0;i<n;i++{
		if lists[i] != nil{
			heap.Push(h, lists[i]) 、// 将所有的链表头节点放入堆中
		}
	}
	if h.Len() == 0{
		return nil
	}

	head := heap.Pop(h).(*ListNode)
	pre := head
	if pre.Next != nil{
		heap.Push(h, pre.Next)
	}
	for h.len() > 0{
		cur := heap.Pop(h).(*ListNode)
		pre.Next = cur
		pre = cur 	
		if cur.Next != nil{
			heap.Push(h, cur.Next)
		}
	}
	return head

}
