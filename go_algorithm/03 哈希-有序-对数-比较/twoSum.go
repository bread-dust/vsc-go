package main

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