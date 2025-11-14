package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("----------------")

	var counter int64
	gs := 100 // gs => go routines
	var wg sync.WaitGroup

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter:", atomic.LoadInt64(&counter))
			time.Sleep(time.Millisecond)
			// runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}


	wg.Wait()
	fmt.Println("----------------")
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)

}