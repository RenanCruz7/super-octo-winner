package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	j := 0
	for j <= 3 {
		fmt.Println(j)
		j++
	}

	for k := range 5 {
		fmt.Println(k)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
