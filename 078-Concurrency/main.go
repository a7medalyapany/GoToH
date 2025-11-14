package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
func main() {

	fmt.Printf("OS\t\t %v\n",runtime.GOOS)
	fmt.Printf("ARCH\t\t %v\n",runtime.GOARCH)
	fmt.Printf("CPUs\t\t %v\n",runtime.NumCPU())
	fmt.Printf("Goroutines\t %v\n",runtime.NumGoroutine())
	fmt.Println("-----------------")

	wg.Add(1)
	go foo()
	bar()
	fmt.Println("-----------------")
	
	wg.Wait()
	fmt.Println("-----------------")
	fmt.Printf("CPUs\t\t %v\n",runtime.NumCPU())
	fmt.Printf("Goroutines\t %v\n",runtime.NumGoroutine())
}

func foo() {
	for i := range 10 {
		fmt.Println("foo", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar", i)
	}
}