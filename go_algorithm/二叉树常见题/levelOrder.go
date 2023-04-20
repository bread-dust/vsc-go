package main

func levelOrder(root *TreeNode)[][]int{
	ans := make([][]int,0)
	if root==nil{
		return ans
	}
	
	queue := make([]TreeNode,0)
	queue = append(queue, *root) // 初始状态加入根节点
	for len(queue) >0{
		// 当前轮，执行tmp次操作
		tmp := len(queue)
		level := make([]int,0)
		for i:=0;i<tmp;i++{
			cur := queue[0]
			if cur.Left !=nil{
				queue = append(queue, *cur.Left)
			}
			if cur.Right !=nil{
				queue = append(queue, *cur.Right)
			}
			level = append(level, cur.Val)
			queue = queue[1:]
		}
		ans = append(ans, level)
	}
	return ans
}
