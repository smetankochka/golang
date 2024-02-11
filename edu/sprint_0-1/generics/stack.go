package main

import (
	"fmt"
)

type Stack[T any] struct {
	elements []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(val T) {
	s.elements = append(s.elements, val)
}

func (s *Stack[T]) Pop() T {
	if len(s.elements) == 0 {
		panic("Stack is empty")
	}
	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return top
}

func main() {
	stack := NewStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack.Pop()) // Output: 3
	fmt.Println(stack.Pop()) // Output: 2
	fmt.Println(stack.Pop()) // Output: 1
}
