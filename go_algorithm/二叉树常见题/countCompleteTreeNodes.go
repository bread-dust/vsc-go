package main

// 当前的头结点是cur，cur在level层	，cur能到达的最深位置
func mostLeftLevel(cur *TreeNode,level int)int{
	for cur!=nil{
		level++
	cur = cur.Left
	}
	return level-1
}

func countNodes(root *TreeNode)int{
	if root ==nil{
		return 0
	}
	return count(root,1,mostLeftLevel(root,1))
}

func count(root *TreeNode,level,h int) int{
	if level==h{
		return 1
	}
	if mostLeftLevel(root.Right,level+1)==h{
		return (1<<(h-level)) + count(root.Right,level,h)
	}else{
		return (1<<(h-level-1)) +count(root.Left,level+1,h)
	}
}





