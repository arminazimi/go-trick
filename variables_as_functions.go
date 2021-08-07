package main

import "fmt"

var funcVar func(int) int

func incfn(x int) int {
	return x + 1
}

func main() {
	funcVar = incfn
	fmt.Println(funcVar(1))
}
