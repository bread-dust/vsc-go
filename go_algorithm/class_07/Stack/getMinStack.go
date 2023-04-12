package stack

type MinStack struct {
	data []int
	min []int
}

func Constructor() MinStack{
	return MinStack{
		data: make([]int, 0),
		min:  make([]int, 0),
	}
}

// Pop
func (ms *MinStack) Pop() {
	n := len(ms.data) -1
	ms.data = ms.data[:n]	
	ms.min = ms.min[:n]
}

// Top
func (ms *MinStack) Top() int {
	return ms.data[len(ms.data)-1]
}

// GetMin
func (ms *MinStack) GetMin()int{
	return ms.min[len(ms.min)-1]
}

// Push 
func (ms *MinStack) Push(x int) {
	ms.data = append(ms.data, x)
	n := len(ms.min)
	if n==0||x<=ms.min[n-1]{ // if x is smaller than the last element in ms.min
		ms.min = append(ms.min, x)
	}else{
		ms.min = append(ms.min,ms.min[n-1])
	}
}