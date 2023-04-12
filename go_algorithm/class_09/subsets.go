package main

// 不重复数组的子集
// htpps://leetcode-cn.com/problems/subsets/

func subsets(nums []int)[][]int{
	ans := make([][]int,0)
	process(nums,0,make([]int,len(nums)),0,&ans)
	return ans
}

// nums[index...] 每个位置的数既可以要也可以不要
// 
func process(nums []int,index int,path []int,size int,ans *[][]int ){
	if index == len(nums){
		cur := make([]int,size)
		for i:=0;i<size;i++{
			cur[i] = path[i]
		}
		*ans = append(*ans,cur) // 一定要加*，否则是值拷贝
	}else{
		// 不要
		process(nums,index+1,path,size,ans)
		path[size] = nums[index]
		// 要
		process(nums,index+1,path,size+1,ans)

	}
}