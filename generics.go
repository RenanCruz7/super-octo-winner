package main

import "fmt"

/*
Generics helps to write flexible and reusable code by allowing you to define functions,
types, and data structures that can work with any data type. This is achieved through the use of type parameters,
which are placeholders for the actual types that will be used when the code is instantiated.

*/

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item
}

func main() {

	// Stack for ints
	intStack := Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	fmt.Println(intStack.Pop()) // 20

	// Stack for strings
	stringStack := Stack[string]{}
	stringStack.Push("Go")
	stringStack.Push("Generics")
	fmt.Println(stringStack.Pop()) // Generics

	// Stack for floats
	floatStack := Stack[float64]{}
	floatStack.Push(3.14)
	fmt.Println(floatStack.Pop()) // 3.14
}
