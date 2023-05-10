package main

import (
	"fmt"
)


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

func maxChunksToSorted2(arr []int)(int,[]int){
	// head 存储栈顶元素
	// head := arr[0]
	// stack 存储栈
	stack := &Stack{}

	// 遍历数组
	// 如果arr[i] >= head 那么入栈
	for _,num := range arr{
		if stack.Len() == 0 || num >= stack.Top(){
			stack.Push(num)
		}else{
			// head = stack.Top()
			// for stack.Top() > num{
			// 	stack.Pop()
			// }
			// stack.Push(head)
		}
	// 如果arr[i] < head 那么出栈head
	// 再将栈顶元素与arr[i]比较
	// 如果栈顶元素大于arr[i] 那么继续出栈
	// 将head入栈 各区间的最大值
	// 返回栈的长度
	
	}
	return stack.Len(),stack.data
}



func main() {
	a := []int{1,1,0,0,1}
	fmt.Println(maxChunksToSorted2(a))

}

