package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	res := add(3, 5)
	fmt.Println("Result:", res)
}
