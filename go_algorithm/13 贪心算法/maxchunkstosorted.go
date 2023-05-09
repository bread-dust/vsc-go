package main

func maxChunksToSorted(arr []int)int{
	n := len(arr)
	min := make([]int,n)
	
	min[n-1] = arr[n-1]
	for i:=n-2;i>=0;i--{
		if arr[i] < min[i+1]{
			min[i] = arr[i]
		}else{
			min[i] = min[i+1]
		}
	}

	ans := 1
	max := arr[0]
	for i:=0;i<n;i++{
		if max <= min[i]{
			ans++
		}
		if arr[i] > max{
			max = arr[i]
		}
	}
	return ans
}

func maxChunksToSorted(arr []int) (ans int) {
    mx := 0
    for i, x := range arr {
        if x > mx {
            mx = x
        }
        if mx == i {
            ans++
        }
    }
    return
}

