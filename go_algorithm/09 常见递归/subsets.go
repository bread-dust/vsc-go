package main

// 不重复数组的子集
// htpps://leetcode-cn.com/problems/subsets/

func subsets(nums []int)[][]int{
	ans := make([][]int,0)
	process(nums,0,make([]int,len(nums)),0,&ans)
	return ans
}

// nums[index...] 每个位置的数既可以要也可以不要
// 当前数来到index 位置，index可要可不要
// nums[0...index-1] 已经做完的决定
// 之前的决定收集了哪些结果，放在path里
// size,path已经使用的长度，选择的数字可能不是n个
// 之前选择的数字都在path[0..size-1]
// ans，收集答案的容器
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
		// 要
		path[size] = nums[index]
		process(nums,index+1,path,size+1,ans)

	}
}

/*
临界条件：下标超出范围
{
	i<子集长度
	tmp[i] = path[i]
	ans = append(ans,tmp)
}

求每个子集的过程
{
	不加自身
	index+1
	path[size] = nums[index]  //将自己加到单个子集的末尾
	size+1,index+1
}
*/


