给定一个数组arr，和一个数num，请把小于等于num的数放在数组的左边，大于num的数放在数组的右边。

思路：
1. 假定一个小于等于区域，初始为空。
2. 遍历数组，
   1. 如果当前元素小于等于num，就把当前元素放到小于等于区域的右边，同时小于等于区域的右边界向右移动一位，
   2. 如果当前元素大于num，就不做任何操作，继续遍历下一个元素。

```go
// arr[l...r],给定pivot 为划分值
func partition(arr []int, l, r, pivot int) {
	//小于等于区域右边界
	lessEqual := l - 1
	for i := l; i <= r; i++ {
		if arr[i] <= pivot {
			lessEqual++
			arr[lessEqual], arr[i] = arr[i], arr[lessEqual]
		}
	}
}

```

