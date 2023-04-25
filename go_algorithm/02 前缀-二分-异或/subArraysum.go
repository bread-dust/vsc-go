func subArraySum1(nums []int,k int)int{
	len := len(nums)
    if len==0{
        return 0
    }

    sum := make([]int,len)
    sum[0]=nums[0]
    right :=0
    for i:=1;i<len;i++{
        sum[i]=sum[i-1]+nums[i]
    }

    for m:=0;m<len;m++{
        for n:=m;n<len;n++{
            if sum[n]-sum[m]+nums[m]==k {
                right++
            }
        }
	}
    return right
}

func subArraySum2(nums []int,k int)int{
	l := len(nums)
    preSum := make(map[int]int)
    preSum[0] = 1
    sum := 0
    res := 0
    for i:=0;i<l;i++{
        sum+=i
        res += preSum[sum-k]
        preSum[sum]++
    }
    return res
}


