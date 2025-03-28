package main

import (
	"fmt"
	"math"
)

const Pi = 3.14
const (
	a = 1
	b = 2
	c = 5 * 5
	d = a + b + c
)

func main() {
	// Constantes podem ser declaraedas em um bloco e em qualquer lugar do código

	fmt.Println("Pi:", Pi)
	fmt.Println("C:", c)
	fmt.Println("D:", d)
	fmt.Println("A:", a)
	fmt.Println("B:", b)
	fmt.Println("Pi from math package:", math.Pi)
	fmt.Println(math.Log(c))
}
