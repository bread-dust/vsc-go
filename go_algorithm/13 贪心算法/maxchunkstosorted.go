package main


// www.leetcode-cn.com/problems/max-chunks-to-make-sorted/
func maxChunksToSorted(arr []int)int{
	// n 数组长度
	// min 存储从右往左的最小值
	n := len(arr)
	min := make([]int,n)
	min[n-1] = arr[n-1]

	// 从右往左遍历
	/// 如果arr[i] < min[i+1] 那么min[i] = arr[i]
	// 否则min[i] = min[i+1]
	for i:=n-2;i>=0;i--{
		if arr[i] < min[i+1]{
			min[i] = arr[i]
		}else{
			min[i] = min[i+1]
		}
	}

	// ans 答案,max 此区间的最大值
	// 从左往右遍历
	// 如果max <= min[i] 那么ans++
	// 如果arr[i] > max 那么max = arr[i]

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

// 建立一个栈
type Stack struct{
	data []int
}
// Top 弹出栈顶元素
func (s *Stack)Top()int{
	if len(s.data) > 0 {
		return s.data[len(s.data)-1]
	}
	return -1 
}

// Push 入栈
func (s *Stack)Push(x int){
	s.data = append(s.data,x)
}

// Pop 出栈
func (s *Stack)Pop()(val int){
	if len(s.data) >0 {
		val = s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return
	}
	
	return -1
}

// Len 栈的长度
func (s *Stack)Len()int{
	return len(s.data)
}

func maxChunksToSorted2(arr []int)int{
	// head 存储栈顶元素
	var head int
	// stack 存储栈
	stack := &Stack{}

	// 遍历数组
	// 如果arr[i] >= head 那么入栈
	for i:=0;i<len(arr)-1;i++{
		if stack.Len()==0 || arr[i] >= head{
			stack.Push(arr[i])
		}else{
			head = stack.Top()
			for stack.Top() > arr[i]{
				stack.Pop()
			}
			stack.Push(head)
		}
	// 如果arr[i] < head 那么出栈head
	// 再将栈顶元素与arr[i]比较
	// 如果栈顶元素大于arr[i] 那么继续出栈
	// 将head入栈 各区间的最大值
	// 返回栈的长度
	
	}
	return stack.Len()
}

