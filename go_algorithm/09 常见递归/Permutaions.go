package main

//不重复数组的全排列
// www.leetcode-cn.com/problems/permutations/
func permute(nums []int) [][]int{
	ans:=make([][]int,0)
	process2(&nums,0,&ans)
	return ans
}

func process2(nums *[]int,index int,ans *[][]int){
	if index==len(*nums){
		cur := make([]int,len(*nums))
		for i:=0;i<len(*nums);i++{
			cur[i] = (*nums)[i]
		}
		*ans = append(*ans,cur)
	}else{
		for i:=index;i<len(*nums);i++{
			swap(nums,i,index)
			process2(nums,index+1,ans)
			swap(nums,i,index)
		}
	}
}

func swap(arr *[]int,i int,j int){
	tmp:=(*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = tmp 

}

/*

当前下标idex,每次与index-len(a)的元素交换完，
index+1 
交换回来(index,i)
*/
