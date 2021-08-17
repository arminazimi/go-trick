package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

type secretAgent struct {
	person
	ltk bool
}

type human interface {
	speak()
}

func (s secretAgent) speak() {
	fmt.Println(s.firstname, s.lastname, s.ltk)
}

func (p person) speak() {
	fmt.Println(p.firstname, p.lastname)
}

// polymorphism
func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println(h.(person).firstname, h.(person).lastname)
	case secretAgent:
		fmt.Println(h.(secretAgent).firstname, h.(secretAgent).lastname, h.(secretAgent).ltk)

	}
}

func main() {
	p := person{"John", "Smith"}
	sa := secretAgent{
		person: person{"James", "Bond"},
		ltk:    true,
	}
	bar(p)
	bar(sa)
}
