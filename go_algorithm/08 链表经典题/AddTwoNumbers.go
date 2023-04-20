package main
// 链表的两数相加
// https://leetcode-cn.com/problems/add-two-numbers/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	len1 := listLength(l1)
	len2 := listLength(l2)
	head := ternary(len1 > len2, l1, l2).(*ListNode) 
	s := ternary(head==l1,l2,l1).(*ListNode)
	l := head
	last := l
	carry := 0 // 进位
	curNum := 0 // 当前数
	for s != nil{ // 遍历短的链表
		curNum = l.Val + s.Val + carry
		l.Val = curNum % 10 // 当前位
		carry = curNum / 10 // 进位
		last  = l
		l = l.Next
		s = s.Next
	}

	for l != nil{ // 遍历长的链表
		curNum = l.Val + carry
		l.Val = curNum % 10
		carry = curNum / 10
		last = l
		l = l.Next
	}

	if carry != 0{ // 最后一位有进位
		last.Next = &ListNode{
			Val:  carry,
		}
	}
	return head
}

func listLength(l *ListNode)int{
	length := 0
	for l != nil{
		length++
		l = l.Next
	}
	return length
}

// 三目运算符
func ternary(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
