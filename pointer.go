package main

import "fmt"

//get value of a address by *
//get adderss &

type person struct {
	name string
}

func main() {
	a := 42
	fmt.Println(a)
	// memory address of a
	fmt.Println(&a)

	b := &a
	fmt.Println(b)
	// value of b
	fmt.Println(*b)

	*b = 43
	fmt.Println(a)

	x := 0
	foo(&x)
	fmt.Println(x)

	p := person{
		name: "gopher",
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)

}

func foo(y *int) {
	*y = 1
}

func changeMe(p *person) {
	// both work
	(*p).name = "golang"
	// p.name = "golang"
}
