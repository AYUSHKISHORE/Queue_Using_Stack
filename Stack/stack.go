package Stack

import "fmt"

type StackInterface interface {
	Push(int)
	Pop() int
	IsEmpty() bool
}

type Stack struct {
	Elements []int
}

func StackInterfaceFunc() StackInterface {
	return &Stack{}
}

func (S *Stack) IsEmpty() bool {
	return len(S.Elements) == 0
}
func (S *Stack) Push(x int) {
	S.Elements = append(S.Elements, x)
}

func (S *Stack) Pop() int {
	//Remove the element
	if S.IsEmpty() {
		fmt.Println("STACK IS EMPTY")
	} else {
		index := len(S.Elements) - 1
		pop := S.Elements[index]
		S.Elements = S.Elements[:index]
		return pop
	}
	return 0
}
