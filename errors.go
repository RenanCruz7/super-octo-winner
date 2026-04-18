package main

import (
	"errors"
)

func check(arg int) (int, error) {
	if arg == 13 {
		return -1, errors.New("unlucky number")
	}
	return arg + 1, nil
}

var ErrToMakeThePass = errors.New("this is too small")

func makeThePass(arg int) (string, error) {
	if arg == 3 {
		return "", ErrToMakeThePass
	}
	return "You Make The Pass", nil
}

func main() {
	if result, err := check(13); err != nil {
		println("Error:", err.Error())
	} else {
		println("Result:", result)
	}

	if result, err := makeThePass(1); err != nil {
		println("Error:", err.Error())
	} else {
		println("Result:", result)
	}
}
