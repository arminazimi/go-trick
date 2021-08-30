package main

import "fmt"

type customErr struct {
	info string
	id   int
}

func (e customErr) Error() string {
	return fmt.Sprintf("the error is __ %v", e.info)
}

func main() {
	c1 := customErr{
		info: "I need more coffee",
		id:   1,
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println("foo_", e, e.(customErr).id)
}

// that is why
// type error interface {
// 	Error() string
// }
