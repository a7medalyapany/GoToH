package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("----------------")

	counter := 0
	gs := 100 // gs => go routines
	var wg sync.WaitGroup

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			time.Sleep(time.Millisecond)
			// runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}


	wg.Wait()
	fmt.Println("----------------")
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)

}