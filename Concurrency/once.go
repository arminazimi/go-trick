package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once

func main() {
	wg.Add(2)
	go foo()
	go foo()
	wg.Wait()
}

func setup() {
	fmt.Println("Setting up")
}

func foo() {
	once.Do(setup)
	fmt.Println("Done")
	wg.Done()
}
