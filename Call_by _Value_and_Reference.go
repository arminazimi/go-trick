package main

import "fmt"

func main() {
	a := 1
	Call_by_Reference(&a)
	fmt.Println(a)

	b := 1
	Call_by_Value(b)
	fmt.Println(b)
}

func Call_by_Reference(y *int) {
	*y = *y + 1
}

func Call_by_Value(y int) {
	y = y + 1
}
