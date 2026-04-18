package main

import (
	"fmt"
	"iter"
)

func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {

	// Traditional iterators
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, number := range numbers {
		fmt.Println(number)
	}
	// Using the custom iterator
	fmt.Println("Fibonacci sequence:")
	for n := range genFib() {
		if n > 50 {
			break
		}
		fmt.Println(n)
	}
}
