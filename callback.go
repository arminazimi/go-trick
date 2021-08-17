package main

import "fmt"

func sum(x ...int) int {
	var sum int
	for _, v := range x {
		sum += v
	}
	return sum
}
func even(f func(x ...int) int, v ...int) int {
	var u []int
	for _, d := range v {
		if d%2 == 0 {
			u = append(u, d)
		}
	}
	return f(u...)
}

func main() {
	z := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	s := sum(z...)
	fmt.Println("all numbers", s)

	se := even(sum, z...)
	fmt.Println("even numbers", se)

}
