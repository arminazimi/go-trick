package main

import (
	"encoding/json"
	"fmt"
	"os"
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
type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
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

	u := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	users := []user{u}

	fmt.Println(users)

	err = json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("We did something wrong and here's the error:", err)
	}

}
