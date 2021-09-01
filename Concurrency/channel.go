package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
		fmt.Println("s")
	}()
	v := <-c
	fmt.Println(v)
	time.Sleep(time.Second * 2)

}
