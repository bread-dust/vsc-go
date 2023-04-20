bsNearRight
有序数组找<=num最右的位置

1. 条件`arr==nil||len(arr)==0`,返回nil
2. 假定左边界为`l=0`，右边界为`r=n-1`，中点`m=0`,结果下标`ans=-1`
3. 循环条件`l<r`
   1. 中点为`m=l+(r-l)>>2`;
   2. 判断`arr[m]<=num`
      1.  true
          1. 将m赋值给ans,ans=m
          2. 左边界右移，l=m+1
      2.  false
          1.  右边界左移，r=m-1
4. 循环结束，返回ans
```go

func moreEqualMostRight(arr []int, val int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	l := 0
	r := len(arr) - 1
	ans := -1
	for l <= r {
		m := l + (r-1)>>1
		if arr[m] <= val {
			ans = m
			l = m + 1
		} else {
			r = m - 1
		}
		return ans
	}
}

```

