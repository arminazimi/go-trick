package main

import (
	"fmt"
	"sync"
)

type counter struct {
	sync.Mutex
	value int
}

func main() {
	var wg sync.WaitGroup

	c := counter{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go inc(&c, &wg)
	}
	wg.Wait()
	fmt.Println(c.value)

}

func inc(c *counter, wg *sync.WaitGroup) {
	c.Lock()
	c.value++
	c.Unlock()
	wg.Done()
}
