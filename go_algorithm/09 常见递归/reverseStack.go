package main

func reverse(stack *Stack){
	if !stack.IsEmpty(){
		bottom :+ popBottom(stack)
		reverse(stack)
		stack.Push(bottom)
	}
}


// 弹出栈底元素
func popBottom(stack *Stack)int{
	if stack.Size==1{
		return stack.Pop()
	}else{
		cur:=stack.Pop()
		ans:=popBottom(stack)
		stack.Push(cur)
		return ans
	}
}

/*
临界条件：stack!=空
先出后进（出的是栈底，放的是栈底）
bottom1=popbottom,取出栈底元素
reverse
ppush(bottom1)


popBottom :abc -> ab 取出c
先出后进（出的是栈底，放的是Pop）
临界条件：只有一个元素，就返回Pop
cur1=Pop 
ans:=popBottom2
push(cur1)
return ans

*/