package main

/*
判断完全二叉树
1. 任何节点有右无左，false
2. 如果左右节点不双全，后序节点都为叶节点,true
*/
func isCompleteTree(root *TreeNode)bool{
	if root==nil{
		return true //空树设定为完全二叉树
	}
	// 是否发现过左右节点不爽全的节点
	leaf := false
	// 按层遍历
	queue := make([]TreeNode,0)
	queue = append(queue,*root)
	for len(queue)>0{
		cur := queue[0]
		if (cur.Right!=nil&&cur.Left==nil )|| 
			(leaf &&(cur.Left!=nil||cur.Right!=nil)){
			return false
		}
		if cur.Left!=nil{
			queue = append(queue, *cur.Left)
		}
		if cur.Right!=nil{
			queue = append(queue, *root)
		}
		if cur.Left == nil || cur.Right==nil{
			leaf = true
		}
	}
		return true

} 