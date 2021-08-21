package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `pass1234`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(s)
	// fmt.Println(bs)

	var loginpass string
	fmt.Println("Enter Your pass: ")
	fmt.Scanf("%s", &loginpass)

	comp(bs, loginpass)

}
func comp(bs []byte, loginpass string) {
	err := bcrypt.CompareHashAndPassword(bs, []byte(loginpass))
	if err != nil {
		fmt.Println("Password does not match")
	} else {
		fmt.Println("Password matches")
	}

}
