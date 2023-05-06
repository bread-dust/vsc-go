package main
// 搜索二叉树


func isValidBST(root *TreeNode)bool{
	list := make([]TreeNode,0)
	inOder(root,&list)
	for i:=1;i<len(list);i++{
		if list[i-1].Val >= list[i].Val{
		return false
		}
	}
	return true
}

func inOder(root *TreeNode,list *[]TreeNode){
	if root == nil{
		return
	}
	inOder(root.Left,list)
	*list=append(*list, *root)
	inOder(root.Right,list)

}