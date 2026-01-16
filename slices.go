package main

import (
	"fmt"
)

func main() {
	// Creating a slice using a slice literal
	var slice1 []int
	slice1 = make([]int, 5) // Slice of length 5
	fmt.Println("Slice 1:", slice1)

	slice1[0] = 10
	slice1[1] = 20
	slice1[2] = 30
	slice1[3] = 40
	slice1[4] = 50
	fmt.Println("Updated Slice 1:", slice1)

	// Creating and initializing a slice in one line
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice 2:", slice2)

	// Appending elements to a slice this makes the slice grow dynamically
	slice1 = append(slice1, 30, 40, 560)
	fmt.Println("Slice 1:", slice1)

	// Copying a slice
	slice3 := make([]int, len(slice1))
	copy(slice3, slice1)
	fmt.Println("Slice 3 (copy of Slice 1):", slice3)
}
