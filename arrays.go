package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println(arr)

	// Initializing an array with values
	b := [5]int{10, 20, 30, 40, 50}
	fmt.Println(b)

	// Initializing an array with a shorthand syntax and inferring the size
	c := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(c)
	fmt.Println(len(c))

	// Initializing an array with specific indices
	// If you specify the index with :, the elements in between will be zeroed
	d := [...]int{1, 9: 10}
	fmt.Println(d)

	var twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
