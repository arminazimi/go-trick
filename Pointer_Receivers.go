package main

import "fmt"

type myint int

func (i myint) double() int {
	return int(i) * 2
}

type point struct {
	x int
	y int
}

func (p *point) offsety(v int) {
	p.y = p.y + v
}

func main() {
	z := point{x: 0, y: 10}
	fmt.Println(z)
	z.offsety(4)
	fmt.Println(z)

}
