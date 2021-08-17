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

// func (r receiver)  identifier(parameters) (return) { code}
func (s secretAgent) speak() {
	fmt.Println(s, s.firstname, s.lastname, s.ltk)
}

func main() {
	sa := secretAgent{
		person: person{"James", "Bond"},
		ltk:    true,
	}
	fmt.Println(sa)
	sa.speak()

}
