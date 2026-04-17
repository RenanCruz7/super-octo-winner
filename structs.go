package main

import "fmt"

type person struct {
	name string
	age  int
}

// constructor function for person struct
func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	// This syntax creates a new struct
	fmt.Println(person{"Bob", 20})
	// You can name the fields when initializing a struct
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Renan"})

	// Age will be 42 using the newPerson Constructor
	fmt.Println(newPerson("Jon"))
}
