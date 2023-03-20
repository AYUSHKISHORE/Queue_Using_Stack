package main

import (
	stack "Queue_Using_Stack/Stack"
	"fmt"
)

var St1 stack.Stack
var St2 stack.Stack

// Function to implement Queue using 2 stack
func main() {
	enqueue(1)
	enqueue(2)
	enqueue(3)
	enqueue(4)
	fmt.Println(St1)
	dequeue()
	dequeue()
}

func enqueue(x int) {
	//After every loop size of stack decreases so find the Size priorly
	Stack1Size := len(St1.Elements)
	for i := 0; i < Stack1Size; i++ {
		St2.Push((St1.Pop()))
	}
	St1.Push(x)
	//After every loop size of stack decreases so find the Size priorly
	Stack2Size := len(St2.Elements)
	for i := 0; i < Stack2Size; i++ {
		St1.Push(St2.Pop())
	}
	// S1 -> 1 , S2->
	// S1 -> 2 , S2 -> 1  ,S1 -> 2, 1
	//S1 -> 3 , S2 ->1 , 2 , S1 ->3 , 2 , 1
}
func dequeue() {
	fmt.Println("Poped element", St1.Pop())
}
