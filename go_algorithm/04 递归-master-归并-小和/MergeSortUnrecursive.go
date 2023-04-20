package main

func sortArrayUnrecursive(arr []int) []int {
    if arr == nil || len(arr) < 2 {
        return arr
    }
    n := len(arr)
    help := make([]int, n) 
    mergeSize := 1 // 步长1,2,4,8,16...
    for mergeSize < n { //步长>n已经排好序
        l := 0 // 左组位置
        for l < n {
            if mergeSize >= n-l { // 右组不够了，不用合并了
                break
            }
            m := l + mergeSize - 1
            r := m + mergeSize // 右组位置
            if r > n-1 {
                r = n - 1
            }
            merge(arr, help, l, m, r)
            l = r + 1 // 下一组
        }
        if mergeSize > n/2 { // 步长大于一半，已经排好序
            break
        }
        mergeSize <<= 1
    }
    return arr
}