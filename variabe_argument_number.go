package main

import "fmt"

func main() {
	fmt.Println(getMax(1, 2, 3, 4))
	vslice := []int{1, 2, 3, 4}
	fmt.Println(getMax(vslice...))
	fmt.Println(getMax())
}

func getMax(vals ...int) int {
	fmt.Println(vals)
	if vals == nil {
		return 0
	}
	maxV := -1
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}
