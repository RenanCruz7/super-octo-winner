package main

import "fmt"

func main() {
	/*message := make(chan string) // Sintax to create a channel of type string

	go func() { message <- "hello world from goroutine 1" }() // Send a value into the channel using the <- operator

	go func() { message <- "hello world from goroutine 2" }()

	msg1 := <-message
	msg2 := <-message // Receive the value from the channel
	fmt.Println(msg1)
	fmt.Println(msg2) // print the value*/

	message3 := make(chan string)
	message4 := make(chan string)

	go func() { message3 <- "hello world from goroutine 3" }()

	go func() {
		msg := <-message3
		message4 <- msg + "processed by goroutine 4"

	}()

	finalMessage := <-message4
	fmt.Println(finalMessage)

}
