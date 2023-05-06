package main

func isBalanced(root *TreeNode)bool{
	if root==nil{
		return true
	}
	return process3(root).isBalanced
}

type Info struct{
	isBalanced bool
	Height int
}

func process3(root *TreeNode)Info{
	if root==nil{
		return Info{true,0}
	}
	left:= process3(root.Left)
	right:=process3(root.Right)
	height:=max(left.Height,right.isBalanced)+1
	isBalanced := left.isBalanced && right.isBalanced && abs(left.Height-right.Height)<=1
	return Info{isBalanced: isBalanced,Height: height}
}