package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("cpus:", runtime.NumCPU())
	fmt.Println("goroutines:", runtime.NumGoroutine())
	var counter int64
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("counter:", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
