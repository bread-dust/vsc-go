package main

// 先序序列化

import (
	"strconv"
	"strings"
)

func serialize(root *TreeNode) string {
	if root == nil {
		return "#_"
	}

	cur := strconv.Itoa(root.Val) +"_"
	cur += serialize(root.Left)
	cur += serialize(root.Right)
	return cur

}

func reconstruct(str string)*TreeNode{
	queue := strings.Split(str,"_")
	return	process2(&queue)
}

func process2(queue *[]string) *TreeNode{
	cur := (*queue)[0]
	*queue = (*queue)[1:]
	if strings.Compare(cur,"#")==0{
		return nil
	}else{
		num,_:= strconv.Atoi(cur)
		head:=&TreeNode{Val:num}
		head.Left= process2(queue)
		head.Right=process2(queue)
		return head
	}
}
