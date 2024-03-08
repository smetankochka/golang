package main

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	elements []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(key T) {
	s.elements = append(s.elements, key)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.elements) == 0 {
		var a T
		return a, errors.New("empty")
	}
	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return top, nil
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
