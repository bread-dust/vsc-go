package main

import (
	"fmt"
	"linkedlist"
	"stack"
)

func sln(){
	head := linkedlist.InitList()
	head.Next = &linkedlist.SingledListNode{
		Val: 1,
	}
	head.Next.Next = &linkedlist.SingledListNode{
		Val: 2,
	}
	head.PrintSingledLinkedList()
	head = head.ReverseList()
	head.PrintSingledLinkedList()
}

func dln(){
	dhead := linkedlist.InitDoubleList()

	dnode1 := &linkedlist.DoubleListNode{
		Val: 1,
	}

	dnode2 := &linkedlist.DoubleListNode{
		Val: 2,
	}

	dhead.Next = dnode1
	dnode1.Prev = dhead
	dnode1.Next = dnode2
	dnode2.Prev = dnode1

	dhead.PrintDoubleLinkedList()
	dhead = dhead.ReverseDoubleList()
	dhead.PrintDoubleLinkedList()
}


func ask(){
	fmt.Println("Array Stack")
	ask := stack.InitArrayStack()
	fmt.Println(ask.IsEmpty())
	ask.Push(1)
	ask.Push(2)
	ask.Push(3)
	fmt.Println(ask.Peek())
	fmt.Println(ask.Pop())
	fmt.Println(ask.Peek())
	fmt.Println(ask.Size())
}

func lsk(){

}

func main() {
	sln() // Singled Linked List

	dln()  // Double Linked List

	ask() // Array Stack

	lsk() // Linked Stack
	
}
