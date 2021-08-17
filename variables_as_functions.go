package main

import "fmt"

var funcVar func(int) int
var fg func() = func() { fmt.Println("hi") }

func incfn(x int) int {
	return x + 1
}

func main() {
	funcVar = incfn
	fmt.Println(funcVar(1))
	fg()
}
