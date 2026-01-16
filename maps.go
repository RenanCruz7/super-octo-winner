package main

import (
	"fmt"
)

func main() {
	// Creating and initializing a map
	// sometimes called hashes or dicts in other languages
	// Maps store key-value pairs
	// Keys are unique within a map
	// Values can be of any type
	m := make(map[string]int)

	m["apple"] = 5
	m["banana"] = 3
	m["orange"] = 7
	fmt.Println(m)
	fmt.Println(len(m))

	// Accessing values
	v4 := m["orange"]
	fmt.Println(v4)

	// Deleting a key-value pair
	delete(m, "banana")
	fmt.Println(m)

	// Initializing a map with values
	n := map[string]string{"name": "Alice", "city": "Wonderland"}
	fmt.Println(n)

	// Checking if a key exists
	value, ok := n["name"]
	if ok {
		fmt.Println("Name:", value)
	} else {
		fmt.Println("Name not found")
	}

}
