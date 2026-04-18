package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg     int
	message string
}

func (e argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func testing(arg int) (int, error) {
	if arg == 42 {
		return -1, argError{arg, "argument out of range"}
	}
	return arg + 3, nil
}

func main() {
	_, err := testing(42)
	var ae argError
	if errors.As(err, &ae) {
		fmt.Println(ae.message)
		fmt.Println(ae.arg)
	}
}
