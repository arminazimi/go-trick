package main

import "fmt"

func main() {
	defer fmt.Println("bye")
	fmt.Println("Hello")

	// argument of a deferred call are evaluated immediately
	i := 0
	defer fmt.Println(i + 1)
	i++

}
