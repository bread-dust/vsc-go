package main

//省份问题
// http://www.leetcode-cn.com/problems/number-of-provinces/
func findCircleNum(isConnected [][]int)int{
	n := len(isConnected)
	uf :=NewUnionFind(n)
	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
			if isConnected[i][j] == 1{
				uf.Union(i,j)
			}
		}
	}
	return uf.sets
}

type UnionFind struct {
	parent []int
	size []int
	help []int
	sets int
}

func NewUnionFind(n int) *UnionFind{
	parent := make([]int,n)
	size := make([]int,n)
	help := make([]int,n)
	for i:=0;i<n;i++{
		parent[i] = i
		size[i] = 1
		help[i] = i
	}
	return &UnionFind{
		parent:parent,
		size:size,
		help:help,
		sets:n,
	}
}

func (uf *UnionFind) Union(i,j int){
	iFather := uf.Find(i)
	jFather := uf.Find(j)
	if iFather == jFather{
		return
	}
	uf.parent[iFather] = jFather
	uf.size[jFather] += uf.size[iFather]
	uf.sets--
	if uf.size[jFather] > uf.size[iFather]{
		uf.parent[iFather] = jFather
		uf.size[jFather] += uf.size[iFather]
	}else{
		uf.parent[jFather] = iFather
		uf.size[iFather] += uf.size[jFather]
	}
}