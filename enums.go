package main

import "fmt"

type Day int

const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d Day) String() string {
	switch d {
	case Sunday:
		return "Sunday"
	case Monday:
		return "Monday"
	case Tuesday:
		return "Tuesday"
	case Wednesday:
		return "Wednesday"
	case Thursday:
		return "Thursday"
	case Friday:
		return "Friday"
	case Saturday:
		return "Saturday"
	default:
		return "Unknown"
	}
}

func main() {
	// Usando o "enum"
	var today Day = Friday
	fmt.Println("Today is:", today)

	// Iterando pelos dias
	for d := Sunday; d <= Saturday; d++ {
		fmt.Println(d, "-", d.String())
	}
}
