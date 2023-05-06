package main

type Info struct{
	No int //一定不偷头结点的最大分数
	Yes int // 偷头结点的最大分数
}

func process4(root *TreeNode)Info{
	if root==nil{
		return Info{0,0}
	}

	left:=process4(root.Left)
	right:=process4(root.Right)
	No := max(left.No,left.Yes)+max(right.No,left.Yes)
	
	Yes := root.Val + left.No + right.No
	return Info{No,Yes}
}