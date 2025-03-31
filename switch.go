package main

import (
	"fmt"
	"time"
)

func main() {

	i := 3
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	day := time.Now().Weekday()

	switch day {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	num := 2

	switch {
	case num%2 == 0:
		fmt.Println("Number is even")
	case num%2 != 0:
		fmt.Println("Number is odd")
	}
}
