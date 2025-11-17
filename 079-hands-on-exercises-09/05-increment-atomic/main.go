package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("GO CPUs", runtime.NumCPU())
	fmt.Println("GO Routines", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(100)
	var incrementer int64
	for range 100 {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			fmt.Println(atomic.LoadInt64(&incrementer))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("end value of incrementer =", incrementer)

}
