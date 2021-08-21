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

func main() {
	s := `[{"Firstname":"John","Lastname":"Doe","Age":30},{"Firstname":"Jane","Lastname":"Doe","Age":25}]`
	bs := []byte(s)
	// fmt.Println(s)
	// fmt.Println(bs)

	var people []person

	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)

	for i, p := range people {
		fmt.Println("\nPerson",i)
		fmt.Println("Firstname:", p.Firstname)
		fmt.Println("Lastname:", p.Lastname)
		fmt.Println("Age:", p.Age)

	}



}
