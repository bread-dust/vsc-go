package main

func MaxDepth(root *TreeNode)int{
	if root==nil{
		return 0
	}

	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)
	return max(left,right)+1
}
