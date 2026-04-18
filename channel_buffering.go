package main

import "fmt"

func main() {

	messages := make(chan string, 3)

	messages <- "buffered 1"
	messages <- "channel 2"
	messages <- "channel 3"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
