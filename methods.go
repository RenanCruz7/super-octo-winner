package main

import "fmt"

type rect struct {
	width, height int
}

// receive by value - does not modify the original object
func (r rect) area() int {
	return r.width * r.height
}

// receive by pointer - modify the original object
func (r *rect) area2() int {
	return r.width * r.height
}

func (r rect) perimeter() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perimeter:", r.perimeter())
}
