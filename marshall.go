package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Firstname string
	Lastname  string
	Age       int
}

type colorGroups struct {
	Id     int
	Name   string
	Colors []string
}

func main() {
	p1 := person{
		Firstname: "John",
		Lastname:  "Doe",
		Age:       30,
	}
	p2 := person{
		Firstname: "Jane",
		Lastname:  "Doe",
		Age:       25,
	}
	a := []person{p1, p2}
	fmt.Println(a)
	bs, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	group := colorGroups{
		Id:     1,
		Name:   "Red",
		Colors: []string{"red", "green", "blue"},
	}
	fmt.Println(group)
	bs, err = json.Marshal(group)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
