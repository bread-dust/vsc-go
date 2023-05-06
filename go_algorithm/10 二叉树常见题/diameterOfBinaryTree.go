package main

// 直径表示距离，线的数量
func diameterOfBinaryTree(root *TreeNode)int{
	return process5(root).Diameter -1
}

type Info struct{
	Height int
	Diameter int
}

func process5(cur *TreeNode)Info{
	if cur == nil{
		return Info{0,0}
	}

	left := process5(cur.Left)
	right := process5(cur.Right)
	Height := max(left.Height,right.Height)+1
	p1 := max(left.Diameter,right.Diameter)
	p2 := left.Height +right.Height + 1
	return Info{Height,max(p1,p2)}
}