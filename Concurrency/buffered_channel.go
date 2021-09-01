package main

func main() {
	c := make(chan int, 10)

	go func() {
		for i := 0; i < 8; i++ {
			c <- i
		}
		close(c)
	}()

	for i := range c {
		println(i)
	}

}
