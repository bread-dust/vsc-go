package stack

type ArrayStack struct {
	data []int 
}

func InitArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]int, 0),
	}
}

func (as *ArrayStack) IsEmpty()bool{
	return len(as.data) == 0
}

func (as *ArrayStack) Size() int{
	return len(as.data)
}

func (as *ArrayStack) Push(val int){
	as.data = append(as.data, val)
}

func (as *ArrayStack) Pop() int{
	n := len(as.data)-1 // n is the index of the last element
	ans := as.data[n] // ans is the last element
	as.data  = as.data[:n]  // remove the last element
	return ans
}

func (as *ArrayStack) Peek()int{
	n := len(as.data)-1 // n is the index of the last element
	ans := as.data[n] // ans is the last element
	return ans
}