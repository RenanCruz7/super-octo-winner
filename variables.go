package main

import "fmt"

func main() {
	var name, SecondName string = "John", "Doe"

	var num = 10
	var num2 = 20

	var t = true
	var f = false

	// Shorthand way to declare variables and initialize them
	u := "apple"

	fmt.Println(name + " " + SecondName)
	fmt.Println((num + num2) / 2)
	fmt.Println(t)
	fmt.Println(f)
	fmt.Println(u)
}
