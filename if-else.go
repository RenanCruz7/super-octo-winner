package main

import "fmt"

func main() {
	a := 10
	b := 20

	if a > b {
		fmt.Println("A is greater than B")
	} else if a < b {
		fmt.Println("A is less than B")
	} else {
		fmt.Println("A is equal to B")
	}

	c := 10

	if c%2 == 0 {
		fmt.Println("C is even")
	} else {
		fmt.Println("C is odd")
	}

}
