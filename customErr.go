package main

import (
	"errors"
	"fmt"
	"log"
)

type customErr struct {
	info string
	id   int
	err  error
}

func (e customErr) Error() string {
	return fmt.Sprintf("the error is __ %v", e.info)
}

func main() {
	// ex1
	c1 := customErr{
		info: "I need more coffee",
		id:   1,
	}
	foo(c1)

	// ex2
	_, err := bar(0)
	if err != nil {
		log.Println(err)
	}

}

func foo(e error) {
	fmt.Println("foo_", e, e.(customErr).id)
}

func bar(i int) (int, error) {
	if i == 0 {
		e := errors.New("more coffee needed")
		return 0, customErr{"error", 1, e}
	}
	return i, nil
}

// that is why
// type error interface {
// 	Error() string
// }
