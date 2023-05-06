package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归序
func f(head *TreeNode) {
	if head == nil { //根节点
		return
	}
	// 1. 先序遍历
	fmt.Println(head.Val, " ")  
	f(head.Left)              
	// 2. 后序遍历
	fmt.Println(head.Val, " ")  
	f(head.Right)
	// 3. 中序遍历
	fmt.Println(head.Val, " ") 
}
