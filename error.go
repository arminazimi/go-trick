package main

import (
	"errors"
	"log"
)

var Err = errors.New("number is less than 10")

func main() {
	_, err := foo(5)
	if err != nil {
		log.Fatal(err)
	}

}

func foo(x int) (int, error) {
	if x < 10 {
		return 0, Err
	}
	return x, nil
}
