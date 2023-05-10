package main

// 小岛问题
// 递归
// http://www.leetcode-cn.com/problems/number-of-islands/
func numIslands(grid [][]byte)int{
	islands :=0
	n := len(grid)
	m := len(grid[0])
	for i:=0;i<n;i++{
		for j:=0;j<m;j++{
			if grid[i][j] == '1'{
				islands++
				infect(&grid,i,j,n,m)
			}
		}
	}
	return islands
}

func infect(grid *[][]byte,i,j,n,m int){
	if i < 0 || i >= n || j < 0 || j >= m || (*grid)[i][j] != '1'{
		return
	}
	// 不越界且grid[i][j] == '1'
	(*grid)[i][j] = '2'
	infect(grid,i+1,j,n,m)
	infect(grid,i-1,j,n,m)
	infect(grid,i,j+1,n,m)
	infect(grid,i,j-1,n,m)
}

// 并查集
func numIslands2(grid [][]byte)int{
	n:=len(grid)
	m:=len(grid[0])
	uf := NewUnionFind(&grid,n,m)
	for i:=0;i<n;i++{
		for j:=0;j<m;j++{
			if grid[i][j] == '1'{
				if i> 0 && grid[i-1][j] == '1'{ //左边
					uf.Union(i-1,j,i,j,m)
				}
				if j > 0 && grid[i][j-1] == '1'{ //上边
					uf.Union(i,j-1,i,j,m)
				}

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


func NewUnionFind(grid *[][]byte,n,m int) *UnionFind{
	parent := make([]int,n*m)
	size := make([]int,n*m)
	help := make([]int,n*m)
	for i:=0;i<n;i++{
		for j:=0;j<m;j++{
			index := i*m+j
			parent[index] = index
			size[index] = 1
			help[index] = index
		}
	}
	return &UnionFind{
		parent:parent,
		size:size,
		help:help,
		sets:n*m,
	}
}