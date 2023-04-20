package main

// 链表是否回文
// https://leetcode-cn.com/problems/palindrome-linked-list/
// 时间复杂度O(n), 空间复杂度O(N)
func isPalindrome1(head *ListNode) bool{
	stack := &Stack{data:make([]int,0)}
	cur := head
	for cur != nil{
		stack.Push(cur.Val)
		cur= cur.Next
	}
	for head !=nil{
		if head.Val != stack.Pop(){
			return  false
		}
		head = head.Next
	}
	return true
}

// 时间复杂度O(N),空间复杂度O(1)
func isPalindrome2(head *ListNode) bool{

	if head == nil || head.Next == nil{
		return true
	}

	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil{ // 快指针走到尾部
			slow = slow.Next // 慢指针走到中间
			fast = fast.Next.Next // 快指针走到尾部
	}

	// 反转后半部分链表
	fast = slow.Next
	slow.Next = nil
	var pre *ListNode
	for fast != nil{
		pre = fast.Next
		fast.Next = slow
		slow = fast
		fast = pre
	}

	pre = slow // 尾部
	fast = head // 头部
	ans := true

	// 逐个比较
	for slow !=nil && fast !=nil{
		if slow.Val != fast.Val{
			ans = false
			break
		}
		slow = slow.Next
		fast = fast.Next
	}

	// 翻转后半部分
	slow = pre.Next 
	pre.Next = nil
	for slow != nil{
		fast = slow.Next 
		slow.Next = pre
		pre = slow
		slow = fast
	}
	return ans
}