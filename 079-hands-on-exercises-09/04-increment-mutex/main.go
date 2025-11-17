package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("GO CPUs", runtime.NumCPU())
	fmt.Println("GO Routines (start)", runtime.NumGoroutine())

	var wg sync.WaitGroup
	var mu sync.Mutex

	incrementer := 0
	wg.Add(100)

	for range 100 {
		go func() {
			mu.Lock()
			incrementer++
			fmt.Println("incrementer =", incrementer)
			mu.Unlock()

			wg.Done()
		}()
	}

	fmt.Println("GO CPUs", runtime.NumCPU())
	fmt.Println("GO Routines (after launch)", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("end value of incrementer =", incrementer)
}
