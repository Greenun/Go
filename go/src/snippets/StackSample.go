package main

import "fmt"

type Stack struct {
	elements []interface{}
}

func NewStack() *Stack {
	return &Stack{elements: make([]interface{}, 0)}
}

func (stack *Stack) Push(data interface{}) {
	stack.elements = append(stack.elements, data)
}

func (stack *Stack) Pop() interface{} {
	if len(stack.elements) < 1 {
		panic("Stack is Empty")
	}
	var value interface{}
	length := len(stack.elements)
	stack.elements, value = stack.elements[:length-1], stack.elements[length-1]
	return value
}

func (stack *Stack) Size() int {
	return len(stack.elements)
}


func main(){
	var stack *Stack = NewStack()
	fmt.Println(stack.Pop())
	stack.Push(10)
	stack.Push(100)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}