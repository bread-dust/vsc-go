package main

// 到叶节点路径

var has = false

func hasPathSum(root *TreeNode,targetSum int)bool{
	if root ==nil{
		return false
	}
	has = false
	process(root,0,targetSum)
	return has
}

func process(cur *TreeNode,preSum ,targetSum int){
	if cur.Left==nil&&cur.Right==nil{
		has = has || preSum +cur.Val == targetSum
	}else{
		if cur.Left !=nil{
			process(cur.Left,preSum+cur.Val,targetSum)
		}
		if cur.Right!=nil{
			process(cur.Right,preSum+cur.Val,targetSum)
		}
	}
}

