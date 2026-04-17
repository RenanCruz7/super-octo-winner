package main

import (
	"fmt"
	"math"
)

/*
To implement an interface in Go,
we just need to implement all the methods in the interface.
Here we implement geometry on rects
*/
type geometry interface {
	area() float64
	perimeter() float64
}

type rect2 struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect2) area() float64 {
	return r.width * r.height
}

func (r rect2) perimeter() float64 {
	return 2*r.width + 2*r.height
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r := rect2{width: 10, height: 5}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
