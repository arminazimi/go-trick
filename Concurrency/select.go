package main

func main() {

	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(eve, odd, quit)

	receive(eve, odd, quit)
}

func receive(eve, odd, quit <-chan int) {
	for {
		select {
		case v := <-eve:
			println("from eve:", v)
		case v := <-odd:
			println("from odd:", v)
		case v := <-quit:
			println("from quit:", v)
			return
		}
	}
}

func send(eve, odd, quit chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	quit <- 0
}
