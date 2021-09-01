package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		c <- 1
		close(c)
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	v, ok = <-c
	fmt.Println(v, ok)


}
