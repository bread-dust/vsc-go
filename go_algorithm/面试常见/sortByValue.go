package main

import "sort"

// 建立一个键值结构体
type Pair struct{
	key string
	value int
}

// 建立一个pair数组
type Pairs []Pair

// 实现sort接口
func (p Pairs) Len() int {return len(p)}
func (p Pairs) Less(i,j int) bool {return p[i].value < p[j].value}
func (p Pairs) Swap(i,j int) {p[i],p[j] = p[j],p[i]}

func sortMapByValue(m map[string]int) Pairs{
	p := make(Pairs,len(m))
	i:= 0
	for k,v := range m{
		p[i] = Pair{k,v}
		i++
	}
	sort.Sort(p)
	return p	
}