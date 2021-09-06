package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("file/test.txt")
	if err != nil {
		return
	}
	defer file.Close()
	s := []string{"hello", "world"}
	for _, v := range s {
		_, err := fmt.Fprintln(file, v)
		if err != nil {
			fmt.Println(err)
			return
		}

	}

}
