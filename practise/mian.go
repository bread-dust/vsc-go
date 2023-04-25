package main

import (
	"fmt"
)

func main() {
	a:=[]int{1,2,3,4,5,6}
	target := 10
	fmt.Print(TwoSum(a,target))
}

func TwoSum(nums []int,target int)[]int{
	l := len(nums)
	Maps := make(map[int]int,l)
	for i,v := range nums{
		if p,ok:=Maps[target-v] ;ok{
			return []int{i,p}
		}
		Maps[v]=i
	}
	return nil
}