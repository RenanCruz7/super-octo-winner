package main

import "fmt"

/*
range iterates over elements in a variety of built-in data structures.
Let’s see how to use range with some of the data structures we’ve already learned

Here we use range to sum the numbers in a slice. Arrays work like this too
range on arrays and slices provides both the index and value for each entry. If you only want the index, you can ignore the second value by assigning it to _.

range on maps iterates over key/value pairs. The iteration order is not specified and is not guaranteed to be the same from one iteration to the next.
*/

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for i, num := range nums {
		if num == 5 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
}
