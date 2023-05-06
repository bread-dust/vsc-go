package main

// 建立并查集
type UnionFind struct {
	parent []int
	size []int
	help []int
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
	}
}

// i点出发，一直往上找，找到不能再往上的代表点，返回
func (uf *UnionFind) Find(i int) int{
	// 一直往上找
	hi := 0
	for i != uf.parent[i]{
		uf.help[hi] = i
		hi++
		i = uf.parent[i]
	}
	// 变扁平化
	for hi != 0{
		hi--
		uf.parent[uf.help[hi]] = i
	}
	return i
}

func (uf *UnionFind) IsSameSet(x,y int) bool{
	return uf.Find(x) == uf.Find(y)
}

func (uf *UnionFind) Union(x,y int){
	xFather := uf.Find(x)
	yFather := uf.Find(y)
	if x == y{
		return
	}
	// 小的挂在大的下面
	if uf.size[xFather] > uf.size[yFather]{
		uf.parent[yFather] = xFather
		uf.size[xFather] += uf.size[yFather]
	}else{
		uf.parent[xFather] = yFather
		uf.size[yFather] += uf.size[xFather]
	}
}
