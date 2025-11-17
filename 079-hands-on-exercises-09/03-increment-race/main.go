package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("GO CPUs", runtime.NumCPU())
	fmt.Println("GO Routines", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(100)
	incrementer := 0
	for range 100 {
		go func() {
			ic := incrementer
			// time.Sleep(time.Millisecond * 2)
			runtime.Gosched()
			ic++
			incrementer = ic
			fmt.Println("incrementer =", incrementer)
			wg.Done()
		}()
	}

	fmt.Println("GO CPUs", runtime.NumCPU())
	fmt.Println("GO Routines", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("end value of incrementer =", incrementer)

}
