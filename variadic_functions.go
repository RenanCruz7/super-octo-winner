package main

import "fmt"

/*

Variadic functions can be called with any number of trailing arguments.
For example, fmt.Println is a common variadic function.

Here’s a function that will take an arbitrary number of ints as arguments
*/

func sum(nums ...int) {
	fmt.Println(nums)
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	sum(1, 2, 3, 4, 5)

}
