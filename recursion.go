package main

import "fmt"

/*
Go supports recursion, which is the ability of a function to call itself.

Anonymous functions can also be recursive,
but this requires explicitly declaring a variable with var to store the function before it’s defined.


*/

func fact(x int) int {
	if x == 0 {
		return 1
	}
	return x * fact(x-1)
}

func main() {
	fmt.Println(fact(42))

	// declaring a variable to store the function
	var fib func(n int) int

	// call the function
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	// use it
	fmt.Println(fib(42))

}
